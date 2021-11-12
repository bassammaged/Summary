#!/usr/bin/python3.9

'''
    # ------------- Arrays in python
        An array is a collection of items stored at contiguous memory locations. The idea is to store multiple items of the same type together. 
        This makes it easier to calculate the position of each element by simply adding an offset to a base value, i.e., the memory location of 
        the first element of the array (generally denoted by the name of the array).
'''


import array as arr                         # importing array module

def main():

    a   = arr.array('i',[2,3,4,5])         # identify array holds integer values.
    
    # -------- iterate array
    for i in range(0,len(a)):
        print(a[i])   


if __name__ == "__main__":
    main()