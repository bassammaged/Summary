// Book class
import java.io.Serializable;

public class Book implements Serializable{
    
    // attributes for book class.
    int book_id;
    String book_name;

    // Constructor function
    public Book(int id, String name) {
        this.book_id    = id;
        this.book_name  = name;
    }
}

