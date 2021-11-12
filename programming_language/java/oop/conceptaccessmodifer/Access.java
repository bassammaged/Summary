/*
    # ----- What is Access Modifier?
        - You must have seen public, private and protected keywords while practising java programs, these are called access modifiers. An access modifier restricts the 
        access of a class, constructor, data member and method in another class. In java we have four access modifiers:
            1. default:
                - When we do not mention any access modifier, it is called default access modifier. The scope of this modifier is limited to the package only. This means that 
                if we have a class with the default access modifier in a package, only those classes that are in this package can access this class. No other class outside this 
                package can access this class.
            2. private:
                - The scope of private modifier is limited to the class only.
                - Private Data members and methods are only accessible within the class.
                - Note: Class and Interface cannot be declared as private.
                - Note: If a class has private constructor then you cannot create the object of that class from outside of the class.

            3. protected:
                - Protected data member and method are only accessible by the classes of the same package and the subclasses present in any package.
                - Note: Classes cannot be declared protected.
            4. public: 
                - The members, methods and classes that are declared public can be accessed from anywhere.

    # ----- Scope of access modifiers:
            ------------+-------+---------+--------------+--------------+--------
                | Class | Package        | Subclass      | Subclass    |Outside|
                |       |                |(same package)|(diff package)|Class  |
            ————————————+———————+—————————+——————————----+—————————----—+————————
            public      | Yes   |  Yes    |    Yes       |    Yes       |   Yes |    
            ————————————+———————+—————————+—————————----—+—————————----—+————————
            protected   | Yes   |  Yes    |    Yes       |    Yes       |   No  |    
            ————————————+———————+—————————+————————----——+————————----——+————————
            default     | Yes   |  Yes    |    Yes       |    No        |   No  |
            ————————————+———————+—————————+————————----——+————————----——+————————
            private     | Yes   |  No     |    No        |    No        |   No  |
            ------------+-------+---------+--------------+--------------+--------
        
*/

public class Access {
    public static void main(String[] args) {

    }    
}
