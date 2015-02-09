<?php

require_once (dirname(__FILE__). '/../SessionDAO.php');

class MySQLSessionDAO implements SessionDAO {

    private $db;
    
    public function __construct($db) {
        $this->db = $db;
    }

    public function getSession($token) {
        $stmt = $this->db->prepare("SELECT * FROM session ".
                     "WHERE access_token=? AND deleted <> 1 ".
                     "AND unix_timestamp(now())<expire_time LIMIT 1");
        $stmt->execute(array($token));
        $list = $stmt->fetchAll();
        if (count($list) == 0) {
            return null;
        } else {
            return $list[0];
        }
    }
}

?>