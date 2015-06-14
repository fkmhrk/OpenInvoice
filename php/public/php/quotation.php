<?php

require_once (dirname(__FILE__). '/../../libs/model/impl/MySQLConnection.php');
require_once (dirname(__FILE__). '/../../libs/model/impl/MySQLSessionDAO.php');
require_once (dirname(__FILE__). '/../../libs/model/impl/MySQLTradingDAO.php');
require_once (dirname(__FILE__). '/../../libs/model/impl/MySQLTradingItemDAO.php');
require_once (dirname(__FILE__). '/../../libs/model/impl/MySQLCompanyDAO.php');
require_once (dirname(__FILE__). '/../../libs/model/impl/MySQLEnvDAO.php');

require_once (dirname(__FILE__). '/../../libs/view/impl/PDFViewImpl.php');

date_default_timezone_set('Asia/Tokyo');

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
$envDAO = new MySQLEnvDAO($db);

$view = new PDFViewImpl();

// execute
$env = $envDAO->getEnv();

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

// render

$view->writeTitle("御見積書");
$view->writeDate($trading['quotation_date'] / 1000);
$view->writeCompany($company['name'], $title);
$view->writeMyCompany($env);
//$view->writeMyCompany($env['company_name']. "\n". $env['company_address']. "\n". $env['company_tel']);

$summary = $view->writeItemTable(16, 120, $items, $trading['tax_rate']);
$view->writeProduct($trading['product']);
$view->writeTotal("御見積金額計 ￥" . number_format($summary['total']));
$view->output('見積書_'. $company['name']);
?>