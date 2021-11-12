#!C:\Python39

'''
    # ------------- Return statement:
        A return statement is used to end the execution of the function call and “returns” the result (value of the expression following the return keyword) to 
        the caller. The statements after the return statements are not executed. If the return statement is without any expression, then the special value 
        None is returned.

'''

# ------ Function with return value
def get_secret_key(status):
    if status == True:
        return "Secret key: 123544487232"
    else:
        return "Get out of here!"
    print("I decieved you :D")
    

status  = bool(input("Enter the status: "))
ret     = get_secret_key(status)
print(ret)
