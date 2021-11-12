#!/usr/bin/python3

'''
    # -------------- polymorphism:
        The word polymorphism means having many forms. In programming, polymorphism means same function name (but different signatures) being uses for different types.

    # -------------- Polymorphism with class methods:
        Below code shows how python can use two different class types, in the same way. We create a for loop that iterates through a tuple of 
        objects. Then call the methods without being concerned about which class type each object is. We assume that these methods actually exist in each class.

    # -------------- syntax:
        # parent class
        class parent_class:
            attribute(s)
            method(s)

        # child class
        class child_class(parent_class):
            attribute(s)
            method(s)

        # create an object
        child_obj   = child_class()
'''

class bird:
    def intro(self):
        print("There are many types of birds")

    def flight(self):
        print("Most of the birds can fly but some cannot.")

    def __del__(self):
        self.flight()

class sparrow(bird):
    def flight(self):
        print("Sparrows can fly.")

class ostrich(bird): 
  def flight(self): 
    print("Ostriches cannot fly.") 
        
obj_bird = bird() 
obj_spr = sparrow() 
obj_ost = ostrich() 
