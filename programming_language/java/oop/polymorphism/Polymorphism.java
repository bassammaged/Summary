/*
    # ----- What is Polymorphism?
        - olymorphism is one of the OOPs feature that allows us to perform a single action in different ways. For example, lets say we have a class 
        Animal that has a method sound(). Since this is a generic class so we can’t give it a implementation like: Roar, Meow, Oink etc. We had to give a generic message.

    # ----- Why using @override?
        - you describe, @Override creates a compile-time check that a method is being overridden. This is very useful to make sure you do not have a silly signature 
        issue when trying to override. because of overoading technique.
*/

class Animal{
    public void sound(){
        System.out.println("Aminal is making a sound.");
    }
}

class Hourse extends Animal{
    @Override
    public void sound(){
        System.out.println("Neigh");
    }
}

public class Cat extends Animal{
    @Override
    public void sound(){
        System.out.println("Meow");
    }
}


public class Polymorphism {
    public static void main(String[] args) {
        
    } 
    
}
