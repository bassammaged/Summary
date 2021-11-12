#!/usr/bin/python3

'''
    # ---------- Switch Case in Python (Replacement)
        What is the replacement of Switch Case in Python? Unlike every other programming language we have used before, Python does not have a 
        switch or case statement. To get around this fact, we use dictionary mapping.
'''

switcher    = {
    0:"Zero",
    1:"One",
    2:"Two"
}

print(switcher.get(5,"Nothing"))