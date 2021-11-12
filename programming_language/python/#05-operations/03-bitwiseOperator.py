#!/usr/bin/python3

'''
    &	    Bitwise AND	            x & y
    |	    Bitwise OR	            x | y
    ~	    Bitwise NOT	            ~x
    ^	    Bitwise XOR	            x ^ y
    >>	    Bitwise right shift	    x>>
    <<	    Bitwise left shift	    x<<
'''

def main():
    # Examples of Bitwise operators 
    a = 10
    b = 4
    
    # Print bitwise AND operation   
    print(a & b) 
    
    # Print bitwise OR operation 
    print(a | b) 
    
    # Print bitwise NOT operation  
    print(~a) 
    
    # print bitwise XOR operation  
    print(a ^ b) 
    
    # print bitwise right shift operation  
    print(a >> 2) 
    
    # print bitwise left shift operation  
    print(a << 2) 

if __name__ == "__main__":
    main()