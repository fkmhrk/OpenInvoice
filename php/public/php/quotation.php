<?php

define('TAX_RATE', 8);

require_once (dirname(__FILE__). '/../../libs/model/impl/MySQLConnection.php');
require_once (dirname(__FILE__). '/../../libs/model/impl/MySQLSessionDAO.php');
require_once (dirname(__FILE__). '/../../libs/model/impl/MySQLTradingDAO.php');
require_once (dirname(__FILE__). '/../../libs/model/impl/MySQLTradingItemDAO.php');
require_once (dirname(__FILE__). '/../../libs/model/impl/MySQLCompanyDAO.php');

date_default_timezone_set('Asia/Tokyo');

function s($str) {
    return mb_convert_encoding($str, "SJIS", "UTF-8");
}

function getTax($v, $taxRate) {
    return ceil($v * $taxRate / 100);
}

function removeTax($v, $taxRate) {
    return floor($v * 100 / (100 + $taxRate));
}

$db = connect();
if ($db === null) {
    return;
}

$token = $_GET['access_token'];
$tradingId = $_GET['trading_id'];

$sessionDAO = new MySQLSessionDAO($db);
$tradingDAO = new MySQLTradingDAO($db);
$tradingItemDAO = new MySQLTradingItemDAO($db);
$companyDAO = new MySQLCompanyDAO($db);

// execute
$session = $sessionDAO->getSession($token);
if ($session === null) {
    echo 'Wrong token';
    return;
}
$userId = $session['user_id'];

$trading = $tradingDAO->getById($userId, $tradingId);
if ($trading === null) {
    echo 'Wrogn trading ID';
    return;
}
$title = $trading['title_type'] == 0 ? '御中' : '様';

$items = $tradingItemDAO->getListByTradingId($tradingId);
if ($items === null) {
    echo 'Wrogn trading ID';
    return;
}

$company = $companyDAO->getById($trading['company_id']);
if ($company === null) {
    $company = array('name' => '(不明な会社)');
}

//print_r($items);

define('FPDF_FONTPATH', dirname(__FILE__) . '/../../libs/fpdf/font/');
require(dirname(__FILE__).  '/../../libs/fpdf/mbfpdf.php');

$pdf = new MBFPDF('P', 'mm', 'A4');
$pdf->AddMBFont(GOTHIC ,'SJIS');
$pdf->AddMBFont(MINCHO ,'SJIS');
$pdf->SetAutoPageBreak(false);
$pdf->SetMargins(0, 0, 0, 0);

$pdf->AddPage('P');

$pdf->SetFont(MINCHO,'', 18);
$pdf->SetXY(16, 16);
$pdf->write(4, s("御見積書"));

$pdf->SetFont(MINCHO,'', 9);
$pdf->SetXY(180, 16);
$pdf->write(4, s(date('Y年m月d日', $trading['quotation_date'])));

$pdf->SetFont(MINCHO,'', 16);
$pdf->SetXY(16, 32);
$pdf->write(4, s($company['name']. " ". $title));

$pdf->SetFont(MINCHO,'', 12);
$pdf->SetXY(140, 50);
$pdf->MultiCell(54, 8, s("サンプル株式会社\n東京\n03-1111-2222"), 0, 'L', 0);

$pdf->SetFont(MINCHO,'', 12);
$pdf->SetXY(16, 120);
$pdf->setFillColor(192, 192, 192);
$pdf->Cell(10, 8, s('No'), 1, 0, 'C', 1);
$pdf->Cell(70, 8, s('品目'), 1, 0, 'C', 1);
$pdf->Cell(20, 8, s('数量'), 1, 0, 'C', 1);
$pdf->Cell(20, 8, s('単位'), 1, 0, 'C', 1);
$pdf->Cell(30, 8, s('単価'), 1, 0, 'C', 1);
$pdf->Cell(30, 8, s('金額'), 1, 1, 'C', 1);

$sum = 0;
$tax = 0;
$no = 1;
foreach ($items as $item) {
    $price = $item['unit_price'] * $item['amount'];
    if ($item['tax_type'] === 1) {
        $tax += getTax($price, TAX_RATE);
    } else if ($item['tax_type'] === 2) {
        $body = removeTax($price, TAX_RATE);
        $tax += ($price - $body);
    } 

    $sum += $price;
    $pdf->SetX(16);
    $pdf->Cell(10, 8, s('' + $no), 1, 0, 'C', 0);
    $pdf->Cell(70, 8, s($item['subject']), 1, 0, 'L', 0);
    $pdf->Cell(20, 8, s($item['amount']), 1, 0, 'R', 0);
    $pdf->Cell(20, 8, s($item['degree']), 1, 0, 'L', 0);
    $pdf->Cell(30, 8, s('￥'. number_format($item['unit_price'])), 1, 0, 'R', 0);
    $pdf->Cell(30, 8, s('￥'. number_format($price)), 1, 1, 'R', 0);

    ++$no;
}
$total = $sum + $tax;

$pdf->SetX(136);
$pdf->Cell(30, 8, s('小計'), 1, 0, 'R', 1);
$pdf->Cell(30, 8, s('￥'. number_format($sum)), 1, 1, 'R', 0);
$pdf->SetX(136);
$pdf->Cell(30, 8, s('消費税'), 1, 0, 'R', 1);
$pdf->Cell(30, 8, s('￥'. number_format($tax)), 1, 1, 'R', 0);
$pdf->SetX(136);
$pdf->Cell(30, 8, s('合計'), 1, 0, 'R', 1);
$pdf->Cell(30, 8, s('￥'. number_format($total)), 1, 1, 'R', 0);

$y = $pdf->GetY();

$pdf->SetXY(16, $y + 8);
$pdf->Cell(180, 8, s('備考'), 1, 1, 'C', 1);
$pdf->SetX(16);
$pdf->MultiCell(180, 8, s($trading['product']), 1, 1, 'L', 0);

$pdf->SetFont(MINCHO,'', 16);
$pdf->SetXY(16, 100);
$pdf->write(4, s("御見積金額計 ￥" . number_format($total)));



$pdf->Output('quotation.pdf', 'D');
?>