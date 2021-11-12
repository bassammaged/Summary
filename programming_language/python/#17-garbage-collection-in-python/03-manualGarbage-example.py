#!/usr/bin/python3

import gc 
i   = 0


def create_cycle():
    x       = { }
    x[i+1]  = x
    print(x)

def main():
    global i 
    for i in range(10):
        create_cycle()

    collected   = gc.collect()
    print ("Garbage collector: collected {} objects.".format(collected))

if __name__ == "__main__":
    main()