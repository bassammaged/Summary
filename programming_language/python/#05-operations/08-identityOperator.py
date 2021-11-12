#!/usr/bin/python3

'''
    is      Evaluates to true if the variables on either side of the operator point to the same object and false otherwise.
    
    is not  Evaluates to false if the variables on either side of the operator point to the same object and true otherwise.
'''

def main():
    x = 24
    y    = [10, 20, 30, 40, 50 ]
    list = [10, 20, 30, 40, 50 ] 
    
    if ( x is list ): 
        print("x is NOT present in given list") 
    else: 
        print("x is  present in given list") 
    
    if ( y is not list ): 
        print("y is present in given list") 
    else: 
        print("y is NOT present in given list") 

if __name__ == "__main__":
    main()