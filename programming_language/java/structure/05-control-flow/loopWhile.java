/*
    # ----- Loop:
        loops are used to execute a set of statements repeatedly until a particular condition is satisfied.

    # ----- Syntax:
        while(condition){
            statement(s);
            // incremental/decremental Or Controller
        }
*/

public class loopWhile{
    public static void main(String[] args){
        int i = 30;
        while (i<37){
            System.out.println("I Value: "+i);
            i++;    // incremental
        }
    }
}