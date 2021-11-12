/*
    # ----- What is oop?
        OOP concepts in Java are the main ideas behind Java’s Object Oriented Programming. They are an abstraction, encapsulation, inheritance, and polymorphism. Grasping 
        them is key to understanding how Java works. Basically, Java OOP concepts let us create working methods and variables, then re-use all or part of them without 
        compromising security.

    # ----- Component of class:
        - Attributes:
            Object attributes is the data bundled in an instance of a class. The object attributes are called instance variables or member fields. An instance 
            variable is a variable defined in a class, for which each object in the class has a separate copy.
        - Methods:
            A method is a block of code which only runs when it is called. You can pass data, known as parameters, into a method. Methods are used to perform 
            certain actions, and they are also known as functions.

    # ----- List of OOP Concepts in Java:
        - Abstraction.      
        - Encapsulation.   
        - Inheritance.
        - Polymorphism.

    # ----- Syntax
        class <class_name> {
            attribute(s);
            method(s);
        }

        // create an abject
        <class_name> <obj_name> = new <class_name>();

        // calling attribute(s)
        <obj_name>.attr_name;

        // calling method(s)
        <obj_name>.method_name();
*/

class person {
        // ----- attributes
        int age;
        String name;
    
        // ----- Methods
        public void hello() {
            System.out.println("Welcome From Hello");
        }
        public void setNewName(){
            this.name = "Junior";
        }
}

public class basicClass {
    public static void main(String[] args) {
        person p1       = new person();
        person p2       = new person();
        // set the attributes
        p1.name     = "Moscu";
        p1.age      = 26;

        p2.name     = "Abdallah";
        p2.age      = 23;

        p2.setNewName();
        
        // calling attributes/methods
        System.out.println("Name:" + p1.name);

    }
} 