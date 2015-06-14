<?php

require_once (dirname(__FILE__). '/../PDFView.php');

define('FPDF_FONTPATH', dirname(__FILE__) . '/../../fpdf/font/');
require_once (dirname(__FILE__).  '/../../fpdf/mbfpdf.php');

function s($str) {
    return mb_convert_encoding($str, "SJIS", "UTF-8");
}

function getTax($v, $taxRate) {
    return ceil($v * $taxRate / 100);
}

function removeTax($v, $taxRate) {
    return floor($v * 100 / (100 + $taxRate));
}

class TQPDFViewImpl implements PDFView {
    private $pdf;
    
    private $startX = 20;
    private $accColorR = 195;
    private $accColorG = 13;
    private $accColorB = 35;

    public function __construct() {
        $this->pdf = new MBFPDF('P', 'mm', 'A4');
        $this->pdf->AddMBFont(GOTHIC ,'SJIS');
        $this->pdf->AddMBFont(MINCHO ,'SJIS');
        $this->pdf->AddMBFont(HIRAKAKU_W6 ,'SJIS');
        $this->pdf->SetAutoPageBreak(false);
        $this->pdf->SetMargins(0, 0, 0, 0);
        $this->pdf->AddPage('P');
    }

    public function writeTitle($title) {
        $this->pdf->SetFillColor($this->accColorR, $this->accColorG, $this->accColorB);
        $this->pdf->Rect($this->startX, 10.0, 1.0, 33.0, F);

        $this->pdf->SetFont(HIRAKAKU_W6, '', 14);
        $this->pdf->SetTextColor($this->accColorR, $this->accColorG, $this->accColorB);
        $this->pdf->SetXY(27, 38);
        $this->pdf->write(4, s($title));        
    }

    public function writeDate($dateTime) {
        $this->pdf->SetFont(GOTHIC,'', 6.5);
        $this->pdf->SetTextColor(77, 77, 77);
        $this->pdf->SetDrawColor(153, 153, 153);
        $this->pdf->SetXY(162, 10);
        $this->pdf->Cell(28.5, 5.8, s('No.V0123'), B, 2, 'R', 0);
        
        $this->pdf->SetXY(162, 15.8);
        $this->pdf->Cell(28.5, 5.8, s(date('Y年m月d日', $dateTime)), B, 2, 'R', 0);
    }

    public function writeCompany($name, $title) {
        $this->pdf->SetFont(HIRAKAKU_W6,'', 10);
        $this->pdf->SetTextColor(0, 0, 0);
        $this->pdf->SetDrawColor(51, 51, 51);
        $this->pdf->SetXY(19.4, 59);
        $this->pdf->Cell(81.8, 8.35, s($name. " ". $title), B, 2, 'L', 0);
    }

    public function writeMyCompany($info) {
        $this->pdf->SetFont(HIRAKAKU_W6, '', 8);
        $this->pdf->SetXY(140.5, 54);
        $this->pdf->write(6, s($info['company_name']));
        
        $this->pdf->SetFont(GOTHIC,'', 7);
        $this->pdf->SetXY(140.5, 60.5);
        $this->pdf->write(4.3, s('〒'. $info['company_zip']. ' '. $info['company_address']));
        $this->pdf->SetXY(140.5, 64.8);
        $this->pdf->write(4.3, s('担当：サンプル / '. $info['company_tel']));        
    }
    
    public function writeItemTable($x, $y, $items, $taxRate) {
        $h1 = 8;    // 1行目
        $h2 = 3;      // 2行目
        $h3 = 2;      // 余白行
        $hBank = 12;  // 振込先用
        
        $w0 = $startX; // $w0は開始位置をあわせるために空白として生成
        $w1 = 70.5;
        $w2 = 29;
        $w3 = 20.9;
        $w4 = 20.9;
        $w5 = 28.5;
        
        // 余白を定義
        $mS = 2;
        $mM = 5;
        
        // 表のはじまる位置
        $startTableY = 78;
        
        // 見出し ----------------------------------------------------
        $this->pdf->SetFont(GOTHIC,'', 8);
        $this->pdf->SetTextColor($this->accColorR, $this->accColorG, $this->accColorB);
        $this->pdf->SetDrawColor($this->accColorR, $this->accColorG, $this->accColorB);
        $this->pdf->SetXY(0, $this->startTableY);
        $this->pdf->Cell($w0, $h1, s(''), 0, 0, 'C', 0);
        $this->pdf->Cell($w1, $h1, s('項目 / 明細'), B, 0, 'C', 0);
        $this->pdf->Cell($w2, $h1, s('単価'), B, 0, 'C', 0);
        $this->pdf->Cell($w3, $h1, s('数量'), B, 0, 'C', 0);
        $this->pdf->Cell($w4, $h1, s('単位'), B, 0, 'C', 0);
        $this->pdf->Cell($w5, $h1, s('金額'), B, 1, 'C', 0);

        $sum = 0;
        $tax = 0;
        $no = 1;
        for ($i = 0 ; $i < count($items) ; ++$i) {
            $item = $items[$i];
            $price = $item['unit_price'] * $item['amount'];
            if ($item['tax_type'] == 1) {
                $tax += getTax($price, $taxRate);
            } else if ($item['tax_type'] == 2) {
                $body = removeTax($price, $taxRate);
                $tax += ($price - $body);
            } 
            
            $sum += $price;

            if ($i == 0) {
                $this->pdf->SetFont(GOTHIC,'', 8);
                $this->pdf->SetTextColor(0, 0, 0);
                $this->pdf->Cell($w0, $h1, s(''), 0, 0, 'C', 0);
                $this->pdf->Cell($w1, $h1, s($item['subject']. '（2/24〜3/20）'), 0, 0, 'L', 0);
                $this->pdf->Cell($w2, $h1, s('¥'. number_format($item['unit_price'])), 0, 0, 'R', 0);
                $this->pdf->Cell($w3, $h1, s($item['amount']), 0, 0, 'C', 0);
                $this->pdf->Cell($w4, $h1, s($item['degree']), 0, 0, 'C', 0);
                $this->pdf->Cell($w5, $h1, s('¥'. number_format($price)), 0, 1, 'R', 0);
                
                $this->pdf->SetFont(GOTHIC,'', 7);
                $this->pdf->SetTextColor(77, 77, 77);
                $this->pdf->Cell($w0, $h2, s(''), 0, 0, 'C', 0);
                $this->pdf->MultiCell($w1, $h2, s($item['product']), 0, 1, 'L', 0);
                
                //空白行
                $this->pdf->SetDrawColor(153, 153, 153);
                $this->pdf->Cell($w0, $h3, s(''), 0, 0, 'C', 0);
                $this->pdf->Cell($w1+$w2+$w3+$w4+$w5, $h3, s(''), B, 1, 'L', 0);
            } else {
            }
            /*
            $this->pdf->SetX(16);
            $this->pdf->Cell(10, 8, s('' + $no), 1, 0, 'C', 0);
            $this->pdf->Cell(70, 8, s($item['subject']), 1, 0, 'L', 0);
            $this->pdf->Cell(20, 8, s($item['amount']), 1, 0, 'R', 0);
            $this->pdf->Cell(20, 8, s($item['degree']), 1, 0, 'L', 0);
            $this->pdf->Cell(30, 8, s('￥'. number_format($item['unit_price'])), 1, 0, 'R', 0);
            $this->pdf->Cell(30, 8, s('￥'. number_format($price)), 1, 1, 'R', 0);
            */
            ++$no;
        }
        $total = $sum + $tax;
        
        $this->pdf->SetX(136);
        $this->pdf->Cell(30, 8, s('小計'), 1, 0, 'R', 1);
        $this->pdf->Cell(30, 8, s('￥'. number_format($sum)), 1, 1, 'R', 0);
        $this->pdf->SetX(136);
        $this->pdf->Cell(30, 8, s('消費税'), 1, 0, 'R', 1);
        $this->pdf->Cell(30, 8, s('￥'. number_format($tax)), 1, 1, 'R', 0);
        $this->pdf->SetX(136);
        $this->pdf->Cell(30, 8, s('合計'), 1, 0, 'R', 1);
        $this->pdf->Cell(30, 8, s('￥'. number_format($total)), 1, 1, 'R', 0);
        return array('sum' => $sum, 'tax' => $tax, 'total' => $total);
    }

    public function writeProduct($product) {
        $this->pdf->SetXY(16, $this->pdf->GetY() + 8);
        $this->pdf->Cell(180, 8, s('備考'), 1, 1, 'C', 1);
        $this->pdf->SetX(16);
        $this->pdf->MultiCell(180, 8, s($product), 1, 1, 'L', 0);
    }

    public function writeTotal($total) {
        $this->pdf->SetFont(MINCHO,'', 16);
        $this->pdf->SetXY(16, 100);
        $this->pdf->write(4, s($total));
    }

    public function output($name) {
        $this->pdf->Output($name . '.pdf', 'D');
    }
}
?>