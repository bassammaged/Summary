<?php 

//  Creating a class for further purpose.
class Book {
    public $book_id;
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
echo '[+] Serialized object: ' . $ser_obj . "\n";


// take the serialized object from the user
$want_to_deser  = readline('Enter a serialized object: ');


// deserialized the object
$o          = unserialize($want_to_deser);
echo "[+] Deserialized object: \n";
var_dump($o);
?>