/*
    # ------ Declare variable:
        - To declare a variable follow this syntax:
            data_type variable_name = value;

        # ----- Variables Types:
            - byte          stores integers (whole numbers), without decimals, hold whole number between -128 and 127
            - short         stores integers (whole numbers), without decimals, Its range is -32,768 to 32767.
            - int           stores integers (whole numbers), without decimals, its range is -2,147,483,648 to 2,147,483,647
            - long          stores integers (whole numbers), without decimals, ranging from -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807.

            - float         stores floating point numbers, with decimals, such as 19.99 or -19.99
            - double        stores floating point numbers, with decimals, such as 19.99 or -19.99

            - String        stores text, such as "Hello". String values are surrounded by double quotes
            - char          stores single characters, such as 'a' or 'B'. Char values are surrounded by single quotes
            - boolean       stores values with two states: true or false

            - <type>[]      stores array of same attributes, for example: int[] numbers = [1,2,3,4,5];

        # ----- Variables naming convention in java
            - Variables naming cannot contain white spaces, for example: int num ber = 100; is invalid because the variable name has space in it.
            - Variable name can begin with special characters such as $ and _
            - As per the java coding standards the variable name should begin with a lower case letter, for example int number; For lengthy 
            variables names that has more than one words do it like this: int smallNumber; int bigNumber; (start the second word with capital letter).
            - Variable names are case sensitive in Java.

    # ----- Types of Variables in Java:
        1. Local variable.
            - These variables are declared inside method of the class. Their scope is limited to the method which means that You can’t change their values and 
            access them outside of the method.
        2. Static (or class) variable.
            - Static variables are also known as class variable because they are associated with the class and common for all the instances of class. For example, 
            If I create three objects of a class and access this static variable, it would be common for all.
        3. Instance variable.
            - Each instance(objects) of class has its own copy of instance variable. Unlike static variable, instance variables have their own separate 
            copy of instance variable.

*/

public class Variables {
    
    public static String pi ="3.14";    // Static Variable
    public String type      = "Russia"; // Instance Variable
    
    public static void main(String[] args) {
        int number      = 2000;
        // Create instance/object for Variables class.
        Variables obj   = new Variables();
        Variables obj2  = new Variables();
        // output
        System.out.println(obj.pi);
        System.out.println(obj2.type);
        System.out.println(number);
    }    
}
