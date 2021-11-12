/*
    # ----- do-while loop:
        do-while loop is similar to while loop, however there is a difference between them: In while loop, condition is evaluated before the execution of loop’s body 
        but in do-while loop condition is evaluated after the execution of loop’s body.

    # ----- Syntax:
        do {
            statement(s);
        } while(condition);
*/
public class loopDoWhile{
    public static void main(String[] args) {
        int i = 10;
        do {
            System.out.println(i);
            i++;
        } while(i<9);
    }
}