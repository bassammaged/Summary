#!C:\Python39

'''
    # ------------- Function, what is that?
        A function is a set of statements that take inputs, do some specific computation and produces output. The idea is to put some commonly or repeatedly done task 
        together and make a function, so that instead of writing the same code again and again for different inputs, we can call the function.
        Python provides built-in functions like print(), etc. but we can also create your own functions. These functions are called user-defined functions.

    # ------------- syntax:
        def fun_name(argurments):
            statement(s)
            return/yield value

'''

# ------ Basic function
def even_odd(x):
    if (x % 2 == 0):
        print("Even!")
    else:
        print("Odd!")



num = int(input("Please Enter a number: "))
even_odd(num)

status  = bool(input("Enter the status: "))
ret     = get_secret_key(status)
print(ret)


