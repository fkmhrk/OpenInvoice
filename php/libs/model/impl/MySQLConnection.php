<?php
define('DB_HOST', "db");
define('DB_USER', "openinvoice_test");
define('DB_PASS', "openinvoice");
define('DB_NAME', "openinvoice");

function connect() {
    try {
        $dsn = "mysql:host=" . DB_HOST.
            ";dbname=" . DB_NAME.
            ";charset=utf8";
        $db = new PDO($dsn,
                      DB_USER,
                      DB_PASS,
                      array(PDO::ATTR_EMULATE_PREPARES => false));
        $db->query("SET NAMES utf8");
        return $db;
    } catch (PDOException $e) {
        echo $e;
        return null;
    }
}
?>