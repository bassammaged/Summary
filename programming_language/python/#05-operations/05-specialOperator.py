#!/usr/bin/python3

'''
    is          True if the operands are identical 
    is not      True if the operands are not identical 
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