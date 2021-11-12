/*
    # ----- What is an exception?  
        - Exception handling is one of the most important feature of java programming that allows us to handle the runtime errors caused by exceptions. 

        - An Exception is an unwanted event that interrupts the normal flow of the program. When an exception occurs program execution gets terminated. In such 
        cases we get a system generated error message. The good thing about exceptions is that they can be handled in Java. By handling the exceptions we can 
        provide a meaningful message to the user about the issue rather than a system generated message, which may not be understandable to a user.

    # ----- Why an exception occurs?
        - here can be several reasons that can cause a program to throw exception. For example: Opening a non-existing file in your program, Network connection 
        problem, bad input data provided by user etc.

    # ----- Difference between error and exception:
        - Errors: indicate that something severe enough has gone wrong, the application should crash rather than try to handle the error.
        - Exceptions are events that occurs in the code. A programmer can handle such conditions and take necessary corrective actions. Few examples:
            NullPointerException – When you try to use a reference that points to null.

    # ----- Types of exceptions: 
            - There are two types of exceptions in Java:
                1. Checked exceptions
                2. Unchecked exceptions

    # ----- finally:
        - A finally block contains all the crucial statements that must be executed whether exception occurs or not. 
            
    # ----- Syntax:
        try{
            //statements that may cause an exception
        } catch (exception(type) e(object))‏ {
            //error handling code
        } finally {
            //Statements to be executed
        }

*/

public class Exception {
    public static void main(String args[]) {
        int num1, num2;
        try {
           /* We suspect that this block of statement can throw 
            * exception so we handled it by placing these statements
            * inside try and handled the exception in catch block
            */
           num1 = 0;
           num2 = 62 / num1;
           System.out.println(num2);
           System.out.println("Hey I'm at the end of try block");
        }
        catch (ArithmeticException e) { 
           /* This block will only execute if any Arithmetic exception 
            * occurs in try block
            */
           System.out.println("You should not divide a number by zero");
        }
        catch (Exception e) {
           /* This is a generic Exception handler which means it can handle
            * all the exceptions. This will execute if the exception is not
            * handled by previous catch blocks.
            */
           System.out.println("Exception occurred");
        }  finally{
                System.out.println("This is finally block");
        }  
        System.out.println("I'm out of try-catch block in Java.");
     }
}
