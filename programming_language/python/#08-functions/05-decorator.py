#!/usr/bin/python3

'''
    # --------- what is decorator?
        - A decorator in Python is any callable Python object that is used to modify a function or a class, a reference 
        to a function "func" or a class "C" is passed to a decorator and the decorator returns a modified function or class.

        - Decorators are very powerful and useful tool in Python since it allows programmers to modify the behavior of function or class. Decorators 
        allow us to wrap another function in order to extend the behavior of wrapped function, without permanently modifying it.

    # --------- Syntax for Decorator:
        # 1st way: using colsure technique
        def decorator_func(function_name):

            def inner_func():
                statement(s)
                function_name()
                statement(s)

            return inner_func
        
        def real_func():
            statement(s)

        final_func  = decorator_func(real_func)
        final_func()


        # 2nd way: using decorator sign @
        def decraotor_func(func_name):
            def inner_func():
                statement(s)
                func_name()
                statement(s)
            return inner_func

        @decorator_func
        def real_func():
            statement(s)

        real_func()
'''


# 1st way: using colsure technique
def hello_decorator(func): 
    # inner1 is a Wrapper function in  
    # which the argument is called 
      
    # inner function can access the outer local 
    # functions like in this case "func" 
    def inner1(): 
        print("Hello, this is before function execution") 
  
        # calling the actual function now, inside the wrapper function. 
        func() 
  
        print("This is after function execution") 
          
    return inner1 

def function_to_be_used(): 
    print("This is inside the function !!") 


# 2nd way: using decorator sign @
def decorator_func(func_name):
    def inner_func():
        print("Modified Functions through @ decorator")
        func_name()

    return inner_func

def real_func(text):
    print("real_func: {}".format(text))


def main():
    # 1st way: using colsure technique
    # passing 'function_to_be_used' inside the decorator to control its behavior 
    function_to_be_used_modified = hello_decorator(function_to_be_used) 
    # calling the function 
    function_to_be_used_modified()

    real_func("Hello!")


if __name__ == "__main__":
    main()