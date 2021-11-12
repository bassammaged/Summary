/* 
    # ----- What is super keyword?
        - The super keyword refers to the objects of immediate parent class.
    
        # ----- The use of super keyword:
            1. To access the data members of parent class when both parent and child class have member with same name.
            2. To explicitly call the no-arg and parameterized constructor of parent class.
            3. To access the method of parent class when child class has overridden that method.
*/

class parent{
    public String name = "From the parent";
    public String type;
    
    // parent constructor
    parent(String x){
        this.type   = x;
    }
    
    public String doesIt() {
        return "does It function runs from the parent class";
    }

}

class child extends parent{

    // override the variable
    public String name = "From the child";
    
    child (String y) {
        // pass the parameter to parent constructor.
        super(y);
    }

    // override the function
    public String doesIt(){
        return "Does It function runs from the child class";
    }

    public void GetParent(){
        // calling the parents method and variable.
        System.out.println("[+] Parent variables and methods, calls from the child class");
        System.out.println(super.name);
        System.out.println(super.doesIt());
        System.out.println(super.type);
    }

    public void GetChild(){
        // calling the child method and variable.
        System.out.println("[+] Child variables and methods, calls from the child class");
        System.out.println(this.name);
        System.out.println(this.doesIt());
        System.out.println(this.type);
    }

}


public class Super{
    public static void main(String[] args) {
        child obj1  = new child("Hello");

        obj1.GetParent();
        obj1.GetChild();
    }

}