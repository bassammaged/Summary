#!/usr/bin/python3

'''
    # -------------- for in Loop: 
        For loops are used for sequential traversal. For example: traversing a list or string or array etc. In Python, there is no C style for loop, 
        i.e., for (i=0; i<n; i++). There is “for in” loop which is similar to for each loop in other languages. Let us learn how to use for in 
        loop for sequential traversals.

    # -------------- Syntax :
        for iterator_var in sequence:
            statements(s)   
        else:
            statement(s) 

'''

n = 4
for i in range(0, n):
    print(i)
else:
    print("No Break control executed.")



# Iterating over a list
print("List Iteration")
l = ["geeks", "for", "geeks"]
for i in l:
    print(i)



# Iterating over a String
print("\nString Iteration")    
s = "Geeks"
for i in s :
    print(i)