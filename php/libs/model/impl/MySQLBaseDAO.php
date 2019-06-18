<?php
class MySQLBaseDAO {

    const DB_HOST = "db";
    const DB_USER = "openinvoice_test";
    const DB_PASS = "openinvoice";
    const DB_NAME = "openinvoice";
    
    public $db;
    protected $query;
    
    public function connect(){
        $dsn = "mysql:host=" . self::DB_HOST.
            ";dbname=" . self::DB_NAME.
            ";charset=utf8";
        $this->db = new PDO($dsn,
                            self::DB_USER,
                            self::DB_PASS,
                            array(PDO::ATTR_EMULATE_PREPARES => false));
    }
    
    protected function execute($query){
        return mysql_query($query, $this->db);
    }
    
    protected function toArray($ret) {
        if ($ret) {
            $row_num = mysql_num_rows($ret);
            for ($i = 0 ; $i < $row_num ; ++$i) {
                $obj[$i] = mysql_fetch_object($ret);
            }
        } else {
            $obj = null;
        }
        return $obj;
    }
    
    protected function getLength($ret){
        return mysql_num_rows($ret);
    }
    
    protected function fetch($ret){
        return mysql_fetch_object($ret);
    }
    
    public function close(){
        mysql_close($this->db);
    }
}

$m = new MySQLBaseDAO();
$m->connect();
$stmt = $m->db->prepare("SELECT * FROM trading WHERE deleted <> 1");
$stmt->execute();
$list = $stmt->fetchAll();
print_r($list);
?>
