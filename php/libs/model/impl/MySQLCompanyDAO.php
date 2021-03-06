<?php
require_once (dirname(__FILE__). '/../CompanyDAO.php');

class MySQLCompanyDAO implements CompanyDAO {

    private $db;
    
    public function __construct($db) {
        $this->db = $db;
    }

    public function getById($id) {
        $stmt = $this->db->prepare("SELECT * FROM company ".
                                   "WHERE id=? AND deleted <> 1 LIMIT 1");
        $stmt->execute(array($id));
        $list = $stmt->fetchAll();
        if (count($list) == 0) {
            return null;
        } else {
            return $list[0];
        }        
    }
}
?>