#!/usr/bin/python3

'''
    any()       Returns true if any of the items is True. It returns False if empty or all are false. Any can be thought of as a sequence of OR 
                operations on the provided iterables.

    all()       Returns true if all of the items are True (or if the iterable is empty). All can be thought of as a sequence of AND operations on 
                the provided iterables. It also short circuit the execution i.e. stop the execution as soon as the result is known. 
'''

def main():
    # Here all the iterables are True so all 
    # will return True and the same will be printed 
    print (all([True, True, True, True])) 
    
    # Here the method will short-circuit at the  
    # first item (False) and will return False. 
    print (all([False, True, True, False])) 
    
    # This statement will return False, as no 
    # True is found in the iterables 
    print (all([False, False, False])) 

if __name__ == "__main__":
    main()