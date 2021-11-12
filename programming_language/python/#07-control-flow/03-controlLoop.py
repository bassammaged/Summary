#!/usr/bin/python3

'''
    # -------------- Loop Control Statements: 
        Loop control statements change execution from its normal sequence. When execution leaves a scope, all automatic objects 
        that were created in that scope are destroyed. Python supports the following control statements. 


    # -------------- statements:
        - Continue Statement: It returns the control to the beginning of the loop.
        - Break Statement: It brings control out of the loop
        - Pass Statement: We use pass statement to write empty loops. Pass is also used for empty control statement, function and classes. 
   


'''

# Prints all letters except 'e' and 's' 
for letter in 'geeksforgeeks':  
    if letter == 'e' or letter == 's': 
         continue
    print ('Continue: Current Letter :', letter) 
    var = 10




for letter in 'geeksforgeeks':  
    # break the loop as soon it sees 'e'  
    # or 's' 
    if letter == 'e' or letter == 's': 
            break

print ('Break: Current Letter :', letter)


# An empty loop 
for letter in 'geeksforgeeks': 
    pass
print ('Pass: Last Letter :', letter)