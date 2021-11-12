/*
    # ----- What is aggregation?
        - Aggregation is a special form of association. It is a relationship between two classes like association, however its a directional association, which means 
        it is strictly a one way association. It represents a HAS-A relationship.
*/

class Address {
    public int streetNum;
    public String streetName;
    public String city;
    public String state;
    public String country;

    Address(int streetNum, String streetName, String city, String state, String country){
        this.streetNum      = streetNum;
        this.streetName     = streetName;
        this.city           = city;
        this.state          = state;
        this.country        = country;
    }
}

class Student{
    public String studentID;
    public Address studentAddress;

    Student(String sID, Address studAddress){
        this.studentID          = sID;
        this.studentAddress     = studAddress;
    }

    public void getAll(){
        System.out.println("Student ID: " + this.studentID);
        System.out.println("student address: " + this.studentAddress);
    }
}

public class Aggregation{
    public static void main(String args[]) {
        Address address     = new Address(30,"937","Future","Cairo","EG");
        Student obj1        = new Student("123ABC",address);
        obj1.getAll();
    }
}