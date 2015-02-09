<?php

require_once (dirname(__FILE__). '/../TradingDAO.php');

class MySQLTradingDAO implements TradingDAO {

    private $db;
    
    public function __construct($db) {
        $this->db = $db;
    }

    public function getById($userId, $id) {
        $stmt = $this->db->prepare("SELECT * ".
                                   "FROM trading ".
                                   "WHERE id=? AND assignee=? AND deleted <> 1 LIMIT 1");
        $stmt->execute(array($id, $userId));
        $list = $stmt->fetchAll();
        if (count($list) == 0) {
            return null;
        } else {
            return $list[0];
        }
    }
}
?>