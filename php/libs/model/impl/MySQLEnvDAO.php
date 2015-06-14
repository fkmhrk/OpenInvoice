<?php

require_once (dirname(__FILE__). '/../EnvDAO.php');

class MySQLEnvDAO implements EnvDAO {
    private $db;
    
    public function __construct($db) {
        $this->db = $db;
    }

    public function getEnv() {
        $stmt = $this->db->prepare("SELECT id,value FROM env ".
                                   "WHERE deleted <> 1 ");
        $stmt->execute(array($token));
        $list = $stmt->fetchAll();
        $env = array();
        foreach ($list as $item) {
            $env[$item['id']] = $item['value'];
        }
        return $env;
    }
}

?>