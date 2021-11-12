/*
    # ----- What is Encapsulation?
        - Encapsulation simply means binding object state(fields) and behaviour(methods) together. If you are creating class, you are doing encapsulation. In this guide 
        we will see how to do encapsulation in java program.
        
        - The whole idea behind encapsulation is to hide the implementation details from users. If a data member is private it means it can only be accessed 
        within the same class. No outside class can access private data member (variable) of other class.

        - However if we setup public getter and setter methods to update (for example void setSSN(int ssn))and read (for example  int getSSN()) the private 
        data fields then the outside class can access those private data fields via public methods.


*/

class Citizen {
    private int ssn;
    private String name;
    public int age;
    
    Citizen(String n, int s, int a){
        this.ssn    = s;
        this.age    = a;
        this.name   = n;
    }

    public void getSSN() {
        System.out.println(this.name + "'s SSN: " + this.ssn);
    }

    public void getName() {
        System.out.println(this.name + "'s name: " + this.name);
    }

    public void getAge() {
        System.out.println(this.name + "'s age: " + this.age);
    }

    public void setSSN(int s) {
        this.ssn    = s;
    }

    public void setName(String n) {
        this.name    = n;
    }

    public void setAge(int a) {
        this.age    = a;
    }
}

public class Enacap {
    public static void main(String[] args) {
        Citizen obj1    = new Citizen("Moscu",29509, 25);
        obj1.getSSN();
    }
}
