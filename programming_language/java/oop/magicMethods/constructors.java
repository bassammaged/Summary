/*
    # ----- Constructor:
        Constructor is a block of code that initializes the newly created object. A constructor resembles an instance method in java but it’s not a method as it doesn’t 
        have a return type. In short constructor and method are different(More on this at the end of this guide). People often refer constructor as special type of method 
        in Java. Constructor has same name as the class and looks like this in a java code.
*/
class person {
    String name;
    int age;
    boolean isMan;

    // constructor method
    person(String n, int a,boolean i) {
        System.out.println("Hello, The object has been created");
        this.name   = n;
        this.age    = a;
        this.isMan  = i;
    }
}

public class constructors {
    public static void main(String[] args) {
        person p1   = new person("Moscu",26,true);


        // Calling contents
        System.out.println("Name: "+ p1.name + ", Age: " + p1.age + " isMan: " + p1.isMan);
    }
}
