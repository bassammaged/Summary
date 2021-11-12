#!/usr/bin/python3

'''
    # ------------- Manual Garbage Collection
        Invoking the garbage collector manually during the execution of a program can be a good idea on how to handle memory being consumed by reference cycles.
'''

import gc

collected   = gc.collect()

print("Garbage collector: collected {} objects.".format(collected))