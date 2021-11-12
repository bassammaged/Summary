/*
    # ----- what is final keyword:
        - final keyword is used in several different contexts to define an entity which cannot later be changed.

        - final class:
            - A final class cannot be subclassed. This is done for reasons of security and efficiency. Many of the Java standard library classes are final, 
            for example java.lang.System and java.lang.String. 
        
        - final method:
            - A final method can't be overridden by subclasses. This is used to prevent unexpected behavior from a subclass altering a method that may be crucial 
            to the function or consistency of the class.

        - final variable:
            - A final variable can only be initialized once, either via an initializer or an assignment statement. It does not need to be initialized at the point of declaration: 
            this is called a blank final variable.
*/

public class FinalTest {
    public static void main(String[] args) {

    }
}