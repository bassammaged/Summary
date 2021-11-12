/*
    # ----- overloading:
        - Method Overloading is a feature that allows a class to have more than one method having the same name, if their argument lists are different. It is similar to 
        constructor overloading in Java, that allows a class to have more than one constructor having different argument lists.

    # ----- Three ways to overload a method:
         1. Number of parameters.
            - Example:
                add(int, int)
                add(int, int, int)
         2. Data type of parameters.
            - Example:
                add(int, int)
                add(int, float)
         3. Sequence of Data type of parameters.
            - Example:
                add(int, float)
                add(float, int)
*/
class calculator {

    // Different number of parameters
    public void add(int x, int y, int z) {
        System.out.println(x + y + z); 
    }

    public void add(int x, int y) {
        System.out.println(x + y); 
    }


    // Different data dype
    public void subs(int x, int y) {
        System.out.println(x + y); 
    }

    public void subs(int x, float y) {
        System.out.println(x + y); 
    }


    // Different Sequence
    public void multiple(int x, float y) {
        System.out.println(x + y); 
    }

    public void multiple(float y, int x) {
        System.out.println(x + y); 
    }
}

public class Overloading{
    public static void main(String[] args) {
        calculator obj1     = new calculator();
        obj1.add(5,10);
        obj1.add(5,10,20);
    }

}