#!/usr/bin/python3

'''
    # -------------- Inheritance:
        Inheritance is a way of creating a new class for using details of an existing class without modifying it. The newly formed class 
        is a derived class (or child class). Similarly, the existing class is a base class (or parent class).

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

class http:
    domain  = ''
    version = ''

    def set_domain(self,d):
        self.domain = d
    def set_version(self,v):
        self.version = v

class https(http):
    sec_protocol  = ''

    def set_sec_protocol(self,s_p):
        self.sec_protocol = s_p


# create object
x_domain    = https()

# inherited methods
x_domain.set_domain('https://mrx.com')
x_domain.set_version('http/1.1')

# class object
x_domain.set_sec_protocol('TLS/1.3')


# print out the values
print("Domain: {}, HTTP version: {}, Security Protocol: {}".format(x_domain.domain,x_domain.version,x_domain.sec_protocol))
