/*
    # ----- What is method:
        Functions "Encapsulate" a task (they combine many instructions into a single line of code). Most programming languages provide many built in functions 
        that would otherwise require many steps to accomplish, for example computing the square root of a number. In general, we don't care how a function does 
        what it does, only that it "does it"!

    # ----- Syntax:
        <visiability> <return_type> <method_name>(arguments) {
            statement(s);
            return <var_name>   // incase of return type not void
        }

    # ----- V.I Note: static methods:-
        A method that has static keyword is known as static method. In other words, a method that belongs to a class rather than an instance of a class is known 
        as a static method. We can also create a static method by using the keyword static before the method name.
*/

public class Customfunction {
    
    public static void main(String[] args) {
        welcome();
    }

    public static void welcome(){
        System.out.println("Welcome, there are many of them! outside.");
    }

}
