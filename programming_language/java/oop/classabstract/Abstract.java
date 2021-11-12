/*
    # ----- What is abstact?
        - A class that is declared using “abstract” keyword is known as abstract class. It can have abstract methods(methods without body) 
        as well as concrete methods (regular methods with body). A normal class(non-abstract class) cannot have abstract methods.
        - An abstract class can not be instantiated.
        
    # ---- Why we need an abstract class?
        - Lets say we have a class Animal that has a method sound() and the subclasses(see inheritance) of it like Dog, Lion, Horse, Cat etc. Since the animal 
        sound differs from one animal to another, there is no point to implement this method in parent class. This is because every child class must override this 
        method to give its own implementation details, like Lion class will say “Roar” in this method and Dog class will say “Woof”.

*/

abstract class Animal{
    public String color;
    
    // abstract method
    public abstract void sound(); 

    public abstract void getColor(String c);
}

class Dog extends Animal {
    public void sound(){
        System.out.println("Woof!");
    }

    public void getColor(String c) {
        this.color = c;
    }
}

public class Abstract {
    public static void main(String[] args) {

    }
}