/*
    # ----- Break statement:
        Use break statement to come out of the loop instantly. Whenever a break statement is encountered inside a loop, the control directly comes out of loop and the 
        loop gets terminated for rest of the iterations. It is used along with if statement, whenever used inside loop so that the loop gets terminated for a particular 
        condition.

    # ----- Syntax:
        break;
*/

public class breakTheFlow {
    public static void main(String[] args){
        for (int i=0; i<6; i++) {
            if (i == 3){
                break;
            }
            System.out.println(i);

        }
    }
}