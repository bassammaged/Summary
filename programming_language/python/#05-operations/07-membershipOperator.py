#!/usr/bin/python3

'''
    in      The ‘in’ operator is used to check if a value exists in a sequence or not. Evaluates to true if it finds a variable in the 
            specified sequence and false otherwise.
    
    not in  Evaluates to true if it does not finds a variable in the specified sequence and false otherwise.
'''

def main():
    x = 24
    y = 20
    list = [10, 20, 30, 40, 50 ]; 
    
    if ( x not in list ): 
        print("x is NOT present in given list") 
    else: 
        print("x is  present in given list") 
    
    if ( y in list ): 
        print("y is present in given list") 
    else: 
        print("y is NOT present in given list") 

if __name__ == "__main__":
    main()