/*
    # ---- What is interface?
        - Interface looks like a class but it is not a class. An interface can have methods and variables just like the class but the methods declared in interface are by 
        default abstract (only method signature, no body).

*/

interface Animal {
    public void sound();
    public void structure();
}

class Dog implements Animal{
    public String type;
    public String subtype;
    
    Dog(String t, String st){
        this.type = t;
        this.subtype = st;
    }
    public void sound() {
        System.out.println("Woof!");
    }

    public void structure(){
        System.out.println("Dog has 4 legs");
    }
}

public class Interface {
    public static void main(String[] args){
        Dog obj1    = new Dog("Type","SubType");
        obj1.sound();
        obj1.structure();
    } 
}
