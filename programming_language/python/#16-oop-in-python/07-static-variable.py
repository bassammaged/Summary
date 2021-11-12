#!/usr/bin/python3

class Car:
    wheel = 4                               # Class (Static) Variable
    def __init__(self,b,c):
        self.brand          = b             # Instance Variable
        self.car_class      = c             # Instance Variable

    def __del__(self):
        print("Brand: {}".format(self.brand))
        print("Class: {}".format(self.car_class))
        print("Wheel: {}".format(self.wheel))

c1  = Car("Alfa Romeo",2)
c2  = Car("BMW",3)
m3  = Car("Kawazaki",3)

# re-assign static variable to specific instance (object) 
m3.wheel = 2

# re-assign static variable (global effect)
# Car.wheel = 2 
