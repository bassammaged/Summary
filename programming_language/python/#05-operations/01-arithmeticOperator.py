#!/usr/bin/python3

'''
    +	        Addition: adds two operands                                                     x + y
    -	        Subtraction: subtracts two operands	                                            x - y
    *	        Multiplication: multiplies two operands	                                        x * y
    /	        Division (float): divides the first operand by the second	                    x / y
    //	        Division (floor): divides the first operand by the second	                    x // y
    %	        Modulus: returns the remainder when first operand is divided by the second	    x % y
    **	        Power : Returns first raised to power second	                                x ** y
'''

def main():
    a = 9
    b = 4
    
    # Addition of numbers  
    add = a + b  
    
    # Subtraction of numbers  
    sub = a - b  
    
    # Multiplication of number  
    mul = a * b  
    
    # Division(float) of number  
    div1 = a / b  
    
    # Division(floor) of number  
    div2 = a // b  
    
    # Modulo of both number  
    mod = a % b  
    
    # Power 
    p = a ** b 

if __name__ == "__main__":
    main()