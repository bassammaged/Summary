/* 
    # ----- Continue Statement:
        Continue statement is mostly used inside loops. Whenever it is encountered inside a loop, control directly jumps to the beginning of the loop for next iteration, 
        skipping the execution of statements inside loop’s body for the current iteration. 

    # ----- Syntax:
        continue;
*/
public class continueTheFlow {
    public static void main(String[] args) {
        for (int i=0;i<6;i++) {
            if (i == 4) {
                continue;
            }
            System.out.println(i);
        }
    }
}
