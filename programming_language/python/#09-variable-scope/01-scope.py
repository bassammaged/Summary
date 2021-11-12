#!/usr/bin/python3.9


'''
    Types of namespaces:
        - built-in namespace.
        - global namespace.
        - local namespace.
    When Python interpreter runs solely without and user-defined modules, methods, classes, etc. Some functions like print(), id() are always present, 
    these are built in namespaces. When a user creates a module, a global namespace gets created, later creation of local functions creates the local namespace.
'''

s   = "Hello All!"                  # global variable

def main(): 
    t   = "Hello from main()"       # local variable

    def nested_main():
        r   = "Hellof rom nested_main()"    # nested local variable
        
        global s                            # calling/identify global variable
        s   = "Hello All! changed from nested_main()"
    
    nested_main()

    print(s)

if __name__ == "__main__":
    main()