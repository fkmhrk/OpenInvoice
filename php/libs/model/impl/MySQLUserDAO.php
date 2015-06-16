<?php

require_once (dirname(__FILE__). '/../UserDAO.php');

class MySQLUserDAO implements UserDAO {
    private $db;
    
    public function __construct($db) {
        $this->db = $db;
    }

    public function getById($userId) {
        $stmt = $this->db->prepare("SELECT * FROM user ".
                                   "WHERE id=? AND deleted <> 1 LIMIT 1");
        $stmt->execute(array($userId));
        $list = $stmt->fetchAll();
        return (count($list) == 0) ? null : $list[0];
    }
}

?>