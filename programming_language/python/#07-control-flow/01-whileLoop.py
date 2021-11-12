#!/usr/bin/python3

'''
    # -------------- While Loop: 
        In python, while loop is used to execute a block of statements repeatedly until a given a condition is satisfied. And when the 
        condition becomes false, the line immediately after the loop in program is executed.
    # -------------- Syntax :
        while expression:
            statement(s) 
        else:
            statement(s) 
        
        ==> Note: Using else statement with while loops: As discussed above, while loop executes the block until a condition is satisfied. When the 
                  condition becomes false, the statement immediately after the loop is executed.
                  - You Can using while without else

'''

count   = 0
while(count < 3):
    print("Hello Moscu {}".format(count))
    count += 1
else:
    print("Condition is not running anymore.")

