// package structure.10-string-methods;

/* 
    # ----- String methods:
        - variable.length();            returns the length of the string.
        - variable.toLowerCase();       returns the string in lowercase.
        - variable.toUpperCase();       returns the string in uppercase.

        # ----- find and Replace:
            - variable.indexOf(<value);                                         find the index of the value. returns -1 incase the value doesn't exist.
            - variable.replace(<value_to_find>,<value_to replace_with>);        replace value_to_find with value_to replace_with

*/

public class StringMethods {
    
    public static void main(String[] args) {
        // declare the variable
        String name             = "Bassam Maged";

        // return the length of var.
        int varLength           = name.length();

        // convert it to uppercase
        String nameUpperCase    = name.toUpperCase();

        // find letter z
        int FindResult          = name.indexOf('z');


        // Outputs
        System.out.println("[+] The length of the string: " + varLength);
        System.out.println("[+] Convert it to uppercase: " + nameUpperCase);
        
        if (FindResult != -1) {
            System.out.println("[+] Z has been found *_^");
        } else {
            System.out.println("[+] Z hasn't been found :(");
        }

    }
}
