#!/usr/bin/python3.9


# ------- function for Unpacking example
def sum(x,y,z):
    print("Sum result:",x+y+z)

# ------- function for Packing example
def ShowInputs(*packing):
    for i in packing:
        print(i)

def main():
    '''
        # ---------- Unpacking
            We can use * to unpack the list so that all elements of it can be passed as different parameters.
            Note: ** is used for dictionaries
            
        # ---------- Packing
            When we don’t know how many arguments need to be passed to a python function, we can use Packing to pack all arguments in a tuple.
    '''
    
    # ------- call function for Unpacking example
    nums    = [30,10,3]
    sum(*nums)

    # ------- call function for Packing example
    ShowInputs(1,3,4,5,"Moscu")



if __name__ == "__main__":
    main()