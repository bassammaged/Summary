<?php 

//  Creating a class for further purpose.
class Book {
    private $book_id;
    public $book_name;

    public function __construct($id,$name){
        echo "[+] The object has been created.\n";
        $this->book_id      = $id;
        $this->book_name    = $name;
    }

    public function __destruct() {
        echo "[+] The object has been destruct.\n";
    }
}

// creating serialized object
$b1         = new Book(12,"Lupin");
$ser_obj    = serialize($b1);

// print the serialized object
echo '[+] Serialized object: ';
var_Export($ser_obj);
echo "\n";


// Inejct the malicious serialized object
$want_to_deser  = 'O:4:"Book":2:{s:13:"' . "\0" . 'Book' . "\0" . 'book_id";s:5:"1ABC2";s:9:"book_name";s:13:"Moscu Is Here";}';

// deserialized the object
$o          = unserialize($want_to_deser);
echo "[+] Deserialized object: \n";
var_dump($o);
?>