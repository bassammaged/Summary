/*
    # ----- Inheritance in Java:
        The process by which one class acquires the properties(data members) and functionalities(methods) of another class is called inheritance. The aim of inheritance 
        is to provide the reusability of code so that a class has to write only the unique features and rest of the common properties and functionalities can be extended 
        from the another class.

        # ----- Parent Class:
            The class whose properties and functionalities are used(inherited) by another class is known as parent class, super class or Base class.
        # ----- Child Class:
            The class that extends the features of another class is known as child class, sub class or derived class.


    # ----- Syntax:
        class XYZ extends ABC{
        }

        // where XYZ is a child for ABC.

    # ----- Type of Inhertance:
        1. Single Inhertance:
            - Single Inheritance: refers to a child and parent class relationship where a class extends the another class
            - Example: Class A <-------------------- Class B

        2. Multilevel Inhertance:
            - refers to a child and parent class relationship where a class extends the child class. For example class C extends class B and class B extends class A.
            - Example: Class A <-------------------- Class B <-------------------- Class C
            - Example:                             Class A <-----|
                                                                |---------------- Class B
                                                    Class C <--| 
        3. Hybrid inheritance:
            - refers to a child and parent class relationship where more than one classes extends the same class. For example, classes B, C & D extends the same class A.
            - Example:                             Class A <-----|
                                                                |---------------- Class B
                                                               |----------------- Class C 
                                                      
    # ----- Constructors and Inheritance:
        - constructor of sub class is invoked when we create the object of subclass, it by default invokes the default constructor of super class. Hence, in inheritance the 
        objects are constructed top-down. The superclass constructor can be called explicitly using the super keyword.
        - Example:  super(params);

    # ----- Inheritance and Method Overriding:
        - When we declare the same method in child class which is already present in the parent class the this is called method overriding. In this case when 
        we call the method from child class object, the child class version of the method is called. However we can call the parent class method using super keyword.

        - Example: obj1SubClass.methodname();
        - Example: super.methodname();          #inside the subclass.
*/


// ----- parent class
class human {
    // ----- public common attributes 
    public String name;
    public int age;
    public float height;
    public float weight;
    public char gender;

    // ----- constructor
    public human(String n, int a, float h, float w, char g){
        System.out.println("Human class has been inhertinced.");
        this.name   = n;
        this.age    = a;
        this.height = h;
        this.weight = w;
        this.gender = g;
    }

    // ----- Craete custom Method
    public void humanity(){
        System.out.println(this.name + " Welcome to humanity society.");
    }

}

// ----- child class
class maleChild extends human{

    // ----- Constructor
    maleChild(){
        super("Moscu",26,100,200,'f');
        System.out.println("Male Object has been created.");
    }
    // ----- Custom Method
    protected void king(){
        System.out.println("King is here.");
    } 
    
}


// ----- main class
public class inhertance{
    public static void main(String[] args) {
        maleChild obj1  = new maleChild();
    }
}


