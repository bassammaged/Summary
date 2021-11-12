// Deserialize class
import java.io.*;

public class Deserialize{
    public static void main(String args[]) {
        try {
            // Creating stream for reading the object.
            ObjectInputStream in    = new ObjectInputStream(new FileInputStream("store.ser"));
            Book b1     = (Book)in.readObject();

            System.out.println("Book ID: " + b1.book_id+"\n"+"Book name: "+b1.book_name);

        }
        // handling the error that may be occur.
        catch (Exception e) {
            System.out.println(e);
        }
    }
}