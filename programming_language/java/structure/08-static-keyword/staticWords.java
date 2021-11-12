/*
    # ----- Static Class, Block, Methods and variables:
        Static keyword can be used with class, variable, method and block. Static members belong to the class instead of a specific instance, this means if you make a 
        member static, you can access it without object. 

        - Static variable:
            1. It is a variable which belongs to the class and not to object (instance).
            2. Static variables are initialized only once, at the start of the execution. These variables will be initialized first, before the initialization of any instance variables.
            3. A single copy to be shared by all instances of the class.
            4. A static variable can be accessed directly by the class name and doesn’t need any object. for example: classname.variable

        - Static Block:
            1. Static block is used for initializing the static variables.This block gets executed when the class is loaded in the memory.
            2. A class can have multiple Static blocks, which will execute in the same sequence in which they have been written into the program.

        - Static method:
            1. It is a variable which belongs to the class and not to object (instance).
            2. A static method can access only static data. It can not access non-static data (instance variables) unless it has/creates an instance of the class.
            3. A static method can call only other static methods and can not call a non-static method from it unless it has/creates an instance of the class.
            4. A static method can be accessed directly by the class name and doesn’t need any object, for example classname.method().
            5. A static method cannot refer to this or super keywords in anyway.

        - Static class:
            1. You cannot use the static keyword with a class unless it is an inner class.
            2. A static inner class is a nested class which is a static member of the outer class.
            3. It can be accessed without instantiating the outer class, using other static members.
            4. A static nested class does not have access to the instance variables and methods of the outer class.

        Note: main method is static since it must be be accessible for an application to run before any instantiation takes place.
*/

public class staticWords {
    static String name;
    int age;

    static{
        int num = 97;
        String mystr = "Static keyword in Java";
     }

    staticWords(String n, int a){
        name    = n;
        this.age = a;
        rideOff();
    }
    public static void rideOff() {
        System.out.println(name);
    }
    public static void main(String[] args) {
        
        staticWords obj1    = new staticWords("Moscu", 26);

        // ----- Change static variable for the 2nd object, it will be reflected within the 1st object too.
        staticWords obj2    = new staticWords("test", 26);

        // ----- Access static variable directly by the class name.
        System.out.println(staticWords.name);


    }    
}
