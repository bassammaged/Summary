#!/usr/bin/python3

'''
    # ------------- example description:
        Create decorator to calculate time execution for each function
'''

import time

def time_exec_calc(func):
    def inner_func(*args,**kargs):
        start_time  = time.time()
        func(*args,**kargs)
        end_time    = time.time()
        time_exec   = end_time - start_time
        print("Time Execution for {}: {:.6f}".format(func.__name__,time_exec)) 
    return inner_func

@time_exec_calc
def incremental(num):
    for x in range(0,num):
        x = x + 1
    print('x = ',x)

def main():
    incremental(1000000)

if __name__ == "__main__":
    main()