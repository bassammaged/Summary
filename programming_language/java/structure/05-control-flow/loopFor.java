/*
    # ----- For Loop:
        Loops are used to execute a set of statements repeatedly until a particular condition is satisfied. In Java we have three types of basic loops: for, while 
        and do-while. In this tutorial we will learn how to use “for loop” in Java.
    
    # ----- Syntax:
    for(initialization; condition ; increment/decrement) {
            statement(s);
        }
*/
public class loopFor {
    public static void main(String[] args) {
        for (int i = 0; i<7; i++) {
            System.out.println("The value of i:" +i);
        }

        // ----- Enhanced For loop:
        int arr[] = {1,2,3,4,5,6};
        for (int num:arr) {
            System.out.println("Array: "+num);
        }
    }
}