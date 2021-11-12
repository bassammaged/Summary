#!c:\Python390

def main():
    '''
        # ----------- Strings in Python 
            A string is a sequence of characters. It can be declared in python by using double-quotes. Strings are immutable, i.e., they cannot be changed.
    '''
    name            = "Python"                                                  # a string
    # ----------- String Access
    print("name[0]:",name[0])
    print("name[-1]:",name[-1])

    # ----------- String Slicing
    print("name[2:4]:",name[2:4])

    # ----------- Escape Sequencing in Python
    # Initial String 
    String1 = '''I'm a "Geek"'''
    print(String1) 
    
    # Escaping Single Quote  
    String1 = 'I\'m a "Geek"'
    print(String1) 
    
    # Escaping Doule Quotes 
    String1 = "I'm a \"Geek\""
    print(String1) 
    
    # Printing Paths with the use of Escape Sequences 
    String1 = "C:\\Python\\Geeks\\"
    print(String1) 

    # To ignore the escape sequences in a String using r/R as a raw string
    String1 = "This is \x47\x65\x65\x6b\x73 in \x48\x45\x58"
    print("Printing in HEX with the use of Escape Sequences: ") 
    print(String1) 
    
    # Using raw String to ignore Escape Sequences 
    String1 = r"This is \x47\x65\x65\x6b\x73 in \x48\x45\x58"
    print("Printing Raw String in HEX Format: ") 
    print(String1) 


    '''
        # ----------- Lists in Python 
            Lists are one of the most powerful tools in python. They are just like the arrays declared in other languages. But the most powerful 
            thing is that list need not be always homogeneous. A single list can contain strings, integers, as well as objects. Lists can also be used 
            for implementing stacks and queues. Lists are mutable, i.e., they can be altered once declared.
    '''
    Courses         = ['Math','Physics','CompSci','History']                    # declare list for store multiple vales, list is mutable

    '''
        # ----------- Tuples in Python 
            A tuple is a sequence of immutable Python objects. Tuples are just like lists with the exception that tuples cannot be changed 
            once declared. Tuples are usually faster than lists.
    '''
    Months          = ("Jan","Feb","March")                                     # declare tuble for store multiple values, tuple is immutable.




    x               = 3                                                         # a (integer) whole number                   
    f               = 3.1415926                                                 # a floating point number              
    

    
    Students        = {1415155:"Bassam",1415156:"Moscu"}                        # declare dict for store bultiple values.

    Employees       = {"Bassam","Moscu"}                                        # declare set for store multiple values. on the other hand each time call set variable the return will be in unorder.

    Status1         = True                                                      # declare boolean variable.
    Status2         = False


    # ------- printing variables
    sum     = f + f
    print("[+] the result of f+f =", sum, "The type of sum:",type(sum))
    print("[+] the result of name[1:4] = ",name[1:4])
    print("[+] the result of set Courses",Courses, "The type of Courses:",type(Courses))
    print("[+] the result of set Months",Months, "The type of Months:",type(Months))
    print("[+] the result of set Students",Students, "The type of Students:",type(Students))
    print("[+] the result of set Employees",Employees, "The type of Employees:",type(Employees))
    print("[+] the result of set Status1",Status1, "The type of Status1:",type(Status1))


if __name__ == "__main__":
    main()