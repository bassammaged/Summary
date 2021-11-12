// Serialze class
import java.io.*;

public class Serialize{
    public static void main(String args[]) {
        try {
            Book b1     = new Book(1001,"Arsene Lupin");

            // Creating stream and storing the object
            FileOutputStream fout       = new FileOutputStream("store.ser");
            ObjectOutputStream out      = new ObjectOutputStream(fout);
            out.writeObject(b1);
            out.flush();

            // closing the stream
            out.close();
            System.out.println("Object has been serialized, the output saved into store.ser");

        } 
        // Throw the exception incase of error occur.
        catch(Exception e) {
            System.out.println(e);
        }
    }
}