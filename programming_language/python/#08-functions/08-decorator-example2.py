#!/usr/bin/python3

import time

def mem_factorial(func):
    start_time  = time.time()
    mem = {}
    def inner_func(num):
        if num not in mem:
            mem[num] = func(num)
            print("[+] Notification: The factorial for: {} has added to memorized dict: {}.".format(num,mem))
            print("[+] Result: The factorial for: {} is: {}.".format(num,mem[num]))
            end_time    = time.time()
            print("Executed time = {}".format(end_time - start_time))
            return mem[num]
        else:
            print("[+] Notification: The factorial for: {} has been calculated before.".format(num))
            print("[+] Result: The factorial for: {} is: {}.".format(num,mem[num]))
            end_time    = time.time()
            print("Executed time = {}".format(end_time - start_time))
    return inner_func


@mem_factorial
def factorial(num):
    factorial = 1
    if num > 0:
        for i in range(1,num+1):
            factorial = factorial * i
        return factorial
    else:
        pass


def main():
    factorial(100)
    factorial(9)
    factorial(14)
    factorial(100)
if __name__ == "__main__":
    main()