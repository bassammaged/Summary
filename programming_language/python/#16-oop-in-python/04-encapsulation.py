#!/usr/bin/python3

'''
    # -------------- Encapsulation:
        Using OOP in Python, we can restrict access to methods and variables. This prevents data from direct modification which is called encapsulation. In Python, we 
        denote private attributes using underscore as the prefix i.e single _ or double __.

    # -------------- syntax:
        # class
        class class_name:
            # private attribute
            __name  = ''

            def __init__(self,n):
                self.__name = n
'''

class bird:

    def __init__(self,n):
        self.__type = n
        
    def get_type(self):
        return self.__type

# create object
obj_bird = bird("Canary")

# access private attribute through public class method
print("Bird type: {}".format(obj_bird.get_type()))


# V.I Note: We can access the value of hidden attribute by a tricky syntax:
print("[+] Tricky: Bird type: {}".format(obj_bird._bird__type))
