/*
    # ----- How to compile and run the code:
        1. Save the file as FirstJavaProgram.java. You may be wondering why we have named the file as FirstJavaProgram, the thing is that we should always 
        name the file same as the public class name. In our program, the public class name is FirstJavaProgram, that’s why our file name should be FirstJavaProgram.java

        2. In this step, we will compile the program. For this, open command prompt command: javac FirstJavaProgram.java

        3. After compilation the .java file gets translated into the .class file(byte code). Now we can run the program. To run the program, 
        Command: java FirstJavaProgram   
*/

/*
    public class <classname> {
        // code
    }
    
    - craete a class with specific name, the visiability of the class is public in our case.
*/
public class FirstJavaProgram {

    /*
        public static void main(String[] args)  {
            // code
        }

        - lets break it down to understand it:
            - public: This makes the main method public that means that we can call the method from outside the class.
            - static: We do not need to create object for static methods to run. They can run itself (Self Invoke).
            - void: It does not return anything.
            - main: It is the method name. This is the entry point method from which the JVM can run your program.
            - (String[] args): Used for command line arguments that are passed as strings.
                example: java FirstJavaProgram a b c
    */
    public static void main(String[] args) {
        System.out.println("FirstJavaProgram");
    }
}