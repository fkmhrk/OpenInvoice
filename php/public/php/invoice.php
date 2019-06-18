<?php
require_once (dirname(__FILE__). '/../libs/model/impl/MySQLConnection.php');
require_once (dirname(__FILE__). '/../libs/model/impl/MySQLSessionDAO.php');
require_once (dirname(__FILE__). '/../libs/view/impl/TQInvoicePDFViewImpl.php');

date_default_timezone_set('Asia/Tokyo');

function isValidInput($json) {
    if (!is_array($json)) {
        echo "Invalid input";
        return false;
    }
    $keys = array('access_token', 'customer', 'myCompany', 'items');
    foreach ($keys as $key) {
        if (!array_key_exists($key, $json)) {
            echo "Invalid input : $key not found";
            return false;
        }
    }
    $token = $json['access_token'];
    if (!is_string($token)) {
        echo "Invalid token";
        return false;        
    }
    $items = $json['items'];
    if (!is_array($items)) {
        echo "Invalid items";
        return false;
    }
    return true;
}
 
$db = connect();
if ($db === null) {
    echo 'Failed to connect DB';
    return;
}
$sessionDAO = new MySQLSessionDAO($db);

//$json = json_decode(file_get_contents('php://stdin'), TRUE);
$json = json_decode(file_get_contents('php://input'), TRUE);

// input check
if (!isValidInput($json)) {
    return;
}

$token = $json['access_token'];
$customer = $json['customer'];
$myCompany = $json['myCompany'];
$items = $json['items'];


// token check
$session = $sessionDAO->getSession($token);
if ($session === null) {
    echo 'Wrong token';
    return;
}

$view = new TQInvoicePDFViewImpl();

$view->writeToAddr($customer['address'], $customer['name']);
$view->writeItemTitle($json['item_title']);
$view->writeDate($json['date'] / 1000);
$view->writeMyCompany($myCompany['name'], $myCompany['address']);
$view->writeTitleMessage();
$view->writeItems($items);
$view->output('添え状');

?>