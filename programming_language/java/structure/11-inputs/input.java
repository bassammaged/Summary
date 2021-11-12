/*
    # ----- takes inputs from the user:
        - The Scanner class is used to get user input, and it is found in the java.util package.
        - To use the Scanner class, create an object of the class and use any of the available methods found in the Scanner class documentation. 

        # ----- steps:
            1. import the package.
            2. create an object.
            3. start to use the Scanner object. 

*/

// ----- Import Scanner class from util package
import java.util.Scanner;


public class input {

    public static void main(String[] args) {
        String name = "";
        int age     = 0;

        
        
        // ----- create Scanner object
        Scanner sc  = new Scanner(System.in);

        // ----- Read String (name) value from the user.
        System.out.print("Enter your name: ");
        name    = sc.nextLine();
        
        // ----- Read int (age) value from the user. 
        System.out.print("Enter your age: ");
        age = sc.nextInt();

        // ----- Close/destory the object
        sc.close();
        

        // output:
            System.out.println("Your name: " + name + " & you are " + age + " years old.");
    }
    
}
