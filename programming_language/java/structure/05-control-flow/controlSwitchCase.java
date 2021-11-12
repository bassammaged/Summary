/* 
    # ----- Switch Case:
        Switch case statement is used when we have number of options (or choices) and we may need to perform a different task for each choice.

    # ----- Syntax:
        switch (variable or an integer expression)
        {
            case constant:
                Statemet(s);
                break;
                case constant:
                Statemet(s);
                break;
            default:
                Statemet(s);
        }
*/
public class controlSwitchCase {
    public static void main(String[] args) {
        int num      = 10;
        switch (num) {
            case 10:
                System.out.println("Num is equal to 10");
                break;
                case 20:
                System.out.println("Num is equal to 20");
                break;
                case 30:
                System.out.println("Num is equal to 30");
                break;
            default:
                System.out.println("Num is not expected!");
        }
    }
}
