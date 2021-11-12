/*

    # ----- What is final keyword?
        - final variables are nothing but constants. We cannot change the value of a final variable once it is initialized.

*/

class Demo{  

    final int MAX_VALUE=99;
    void myMethod(){  
       MAX_VALUE=101;
    } 
 }

public class Final {
    Demo obj = new  Demo();  
    obj.myMethod();  
}
