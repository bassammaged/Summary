#!/usr/bin/python3

'''
    # ------------- Class attributes (Built-In):
        Every Python class keeps following built-in attributes and they can be accessed using dot operator like any other attribute

    # ------------- syntax:
        object.__attrName__

    # ------------- Attributes:
        __dict__:       Dictionary containing the class's namespace.
        __doc__         Class documentation string or none, if undefined.
        __class__       Class name.
        __module__      Module name in which the class is defined. This attribute is "__main__" in interactive mode.
        __bases__       possibly empty tuple containing the base classes, in the order of their occurrence in the base class list.
    
'''


class human:
    name    = ''
    age     = 0

    def set_name(self,n):
        self.name = n

    def get_name(self):
        print("Name: {}".format(self.name))



# create object
me  = human()

# call method
me.set_name("Moscu")
me.get_name()

# Using built-in class attributes
print(".__class__: {}".format(me.__class__))