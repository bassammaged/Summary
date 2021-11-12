#!C:\Python39

'''
    # ------------- When to use yield instead of return in Python:
        The yield statement suspends function’s execution and sends a value back to the caller, but retains enough state to enable function to 
        resume where it is left off. When resumed, the function continues execution immediately after the last yield run. This allows its code to produce a 
        series of values over time, rather than computing them at once and sending them back like a list.
'''

# ------ Function with return value
def simpleGeneratorFun(): 
    yield 1
    yield 2
    yield 3

print(list(simpleGeneratorFun()))