#!/usr/bin/python3

'''
    # ------------- Built-In Class methods (Magic methods):
        Magic methods in Python are the special methods which add "magic" to your class. Magic methods are not meant to be invoked directly by you, but the 
        invocation happens internally from the class on a certain action. For example, when you add two numbers using the + operator, internally, the __add__() method will be called.

    # ------------- syntax:
        object.__attrName__()

    # ------------- Magic methods:
        __new__(cls,other)          To get called in an object's instantiation. (when you need to control the creation of a new instance.)
        __init__(self,other)        To get called by the __new__ method. (when you need to control initialization of a new instance.) (Constructors)
        __del__(self)               Destructor method.
    
'''


class human:
    name    = ''

    # def __new__(cls):
    #     return object.__new__(cls)

    def __init__(self,n,a):
        self.name   = n
        self.age    = a

    def set_name(self,n):
        self.name = n

    def get_name(self):
        print("Name: {}".format(self.name))



# create object
me  = human("Moscu",25)

me.get_name()
