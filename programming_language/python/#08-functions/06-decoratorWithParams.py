#!/usr/bin/python3

'''
    # --------- How we are using params with decorators
        - 
    # --------- Syntax:
    def decorator_func(*args,**kargs):
        def inner_func(func_name):
            statement(s)
            func_name()
            statement(s)
        
        return inner_func

    @decorator_func(params)
    def real_func():
        statement(s)

'''

# def decraotor_func(*args,**kargs):
#     def inner_func(func_name):
#         print(args)
#         func_name()

#     return inner_func

# @decraotor_func(test="Moscu")
# def real_func():
#     print("Do some action from real func")
    
# def main():
#     real_func()


def decorator(*args, **kwargs): 
    print("Inside decorator") 
      
    def inner(func): 
          
        # code functionality here 
        print("Inside inner function") 
        print("I like", kwargs['like'])  
          
        func() 
          
    # reurning inner function     
    return inner 
  
@decorator(like = "geeksforgeeks") 
def my_func(): 
    print("Inside actual function") 

