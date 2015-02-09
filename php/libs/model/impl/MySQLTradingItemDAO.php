<?php

require_once (dirname(__FILE__). '/../TradingItemDAO.php');

class MySQLTradingItemDAO implements TradingItemDAO {

    private $db;
    
    public function __construct($db) {
        $this->db = $db;
    }

    public function getListByTradingId($tradingId) {
        $stmt = $this->db->prepare("SELECT * ".
                                   "FROM trading_item ".
                                   "WHERE trading_id=? AND deleted <> 1 ORDER BY sort_order ASC");
        $stmt->execute(array($tradingId));
        return $stmt->fetchAll();
    }
    /*    
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
    */
}
?>