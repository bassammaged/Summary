#!C:\Python39

'''
    # ------------- Note: What is nested function?
        A function which is defined inside another function is known as nested function. Nested functions are able to access variables of the enclosing scope.
        In Python, these non-local variables can be accessed only within their scope and not outside their scope.

    # ------------- Python Closures:
        - A Closure is a function object that remembers values in enclosing scopes even if they are not present in memory.
        - It is a record that stores a function together with an environment: a mapping associating each free variable of the function (variables that are used 
        locally, but defined in an enclosing scope) with the value or reference to which the name was bound when the closure was created.
        - A closure—unlike a plain function—allows the function to access those captured variables through the closure’s copies of their values or 
        references, even when the function is invoked outside their scope.

    # ------------- When and why to use Closures?
        - As closures are used as callback functions, they provide some sort of data hiding. This helps us to reduce the use of global variables.
        - When we have few functions in our code, closures prove to be efficient way. But if we need to have many functions, then go for class (OOP).
'''

def outerFunc(text):
    text = text

    def innerFunc():
        print("From InnerFunc: {}".format(text))
   
    # Note we are returning function 
    # WITHOUT parenthesis 
    return innerFunc

def main():
    result  = outerFunc("Hello There!")
    result()

if __name__ == "__main__":
    main()