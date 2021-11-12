#!/usr/bin/python3

'''
    =	Assign value of right side of expression to left side operand	                                                x = y + z
    +=	Add AND: Add right side operand with left side operand and then assign to left operand	                        a+=b     a=a+b
    -=	Subtract AND: Subtract right operand from left operand and then assign to left operand	                        a-=b       a=a-b
    *=	Multiply AND: Multiply right operand with left operand and then assign to left operand	                        a*=b       a=a*b
    /=	Divide AND: Divide left operand with right operand and then assign to left operand	                            a/=b         a=a/b
    %=	Modulus AND: Takes modulus using left and right operands and assign result to left operand	                    a%=b   a=a%b
    //=	Divide(floor) AND: Divide left operand with right operand and then assign the value(floor) to left operand	    a//=b       a=a//b
    **=	Exponent AND: Calculate exponent(raise power) value using operands and assign value to left operand	            a**=b     a=a**b
    &=	Performs Bitwise AND on operands and assign value to left operand	                                            a&=b     a=a&b
    |=	Performs Bitwise OR on operands and assign value to left operand	                                            a|=b         a=a|b
    ^=	Performs Bitwise xOR on operands and assign value to left operand	                                            a^=b       a=a^b
    >>=	Performs Bitwise right shift on operands and assign value to left operand	                                    a>>=b     a=a>>b
    <<=	Performs Bitwise left shift on operands and assign value to left operand	                                    a <<= b   a= a << b
'''

def main():
    # Examples of Identity operators 
    a1 = 3
    b1 = 3
    a2 = 'GeeksforGeeks'
    b2 = 'GeeksforGeeks'
    a3 = [1,2,3] 
    b3 = [1,2,3] 

    print(a1 is not b1) 
    
    
    print(a2 is b2) 
    
    # Output is False, since lists are mutable. 
    print(a3 is b3) 

if __name__ == "__main__":
    main()