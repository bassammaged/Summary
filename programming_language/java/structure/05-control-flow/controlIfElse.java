/* 
    # ----- what is control statement?
        When we need to execute a set of statements based on a condition then we need to use control flow statements. For example, if a number is greater than 
        zero then we want to print Positive Number but if it is less than zero then we want to print Negative Number. In this case we have two print statements 
        in the program, but only one print statement executes at a time based on the input value. We will see how to write such type of conditions in the java 
        program using control statements.

    # ----- Syntax:
        if(condition){
            Statement(s);
        } else if (condition) {

        } else if (condition) {
            Statement(s);
        } else {
            Statement(s);
        }
*/

public class controlIfElse {
    public static void main(String[] args){
        int num1 = 10;
        int num2 = 20;

        if (num1 > num2) {
            System.out.println("Num1 is bigger than Num2");
        } else if (num2 > num1) {
            System.out.println("Num2 is bigger than Num1");
        } else {
            System.out.println("Num1 is equal to Num2");
        }
    }
}