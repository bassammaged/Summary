#!/usr/bin/python3

'''
    # ------------- What is OOP:
        - Object-oriented programming is a programming paradigm that provides a means of structuring programs so that properties and behaviors are bundled into 
        individual objects.
        - For instance, an object could represent a person with properties like a name, age, and address and behaviors such as walking, talking, breathing, and 
        running. Or it could represent an email with properties like a recipient list, subject, and body and behaviors like adding attachments and sending.

    # ------------- Why OOP:
        - Modularity for easier troubleshooting.
        - Reuse of code through inheritance.
        - Flexibility through polymorphism.
        - Effective problem solving.
        - implement the concept of encapsulation 

    # ------------- Syntax:
        # create class
        class class_name:
            attribute(s)
            def func_name(params):
                statement(s)

        # create an object
        object_name = class_name()
'''

# create human class
class human:
    number  = 0
    name    = "no name"
    age     = 0

    def set_name(self,n):
        self.name   = n
    
    def set_age(self,a):
        self.age    = a


# create object
me  = human()

print("[+] name attribute before editing: {}".format(me.name))

# call method inside the object
me.set_name("Moscu")
print("[+] name attribute after the edition: {}".format(me.name))

# change attribute directly
me.number = 20
print("[+] number attribute after the directly edition: {}".format(me.number))
