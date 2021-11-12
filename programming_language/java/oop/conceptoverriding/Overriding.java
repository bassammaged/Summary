/*
    # ----- What is overriding method?
        - Declaring a method in sub class which is already present in parent class is known as method overriding. Overriding is done so that a child class can give its own 
        implementation to a method which is already provided by the parent class. In this case the method in parent class is called overridden method and the method in child class 
        is called overriding method.

    
*/

class Human{
    //Overridden method
    public void eat()
    {
       System.out.println("Human is eating");
    }
 }
 
 class Boy extends Human{
    //Overriding method
    public void eat(){
       System.out.println("Boy is eating");
    }
 }

public class Overriding {
    public static void main(String[] args) {
        Boy obj = new Boy();
        obj.eat();
    }
}
