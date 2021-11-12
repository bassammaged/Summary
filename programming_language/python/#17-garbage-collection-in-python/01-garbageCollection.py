#!/usr/bin/python3


'''
    # ------------- Garbage Collection in Python
        - Python’s memory allocation and deallocation method is automatic. The user does not have to preallocate or deallocate memory similar to using dynamic memory 
        allocation in languages such as C or C++.

        - Python uses two strategies for memory allocation:
            - Reference counting
            - Garbage collection

     # ------------- How Python2.x works:
        Prior to Python version 2.0, the Python interpreter only used reference counting for memory management. Reference counting works by counting the number 
        of times an object is referenced by other objects in the system. When references to an object are removed, the reference count for an object 
        is decremented. When the reference count becomes zero, the object is deallocated.


    # ------------- Ways to make an object eligible for garbage collection
        x = [] 
        x.append(l) 
        x.append(2) 
        
        # delete the list from memory or  
        # assigning object x to None(Null) 
        del x  
        # x = None 
        The reference count for the list created is now two. However, since it cannot be reached from inside Python and cannot possibly be used again, it is 
        considered garbage. In the current version of Python, this list is never freed.

    # ------------- Automatic Garbage Collection of Cycles
        - Because reference cycles take computational work to discover, garbage collection must be a scheduled activity. Python schedules garbage collection 
        based upon a threshold of object allocations and object deallocations. When the number of allocations minus the number of deallocations is 
        greater than the threshold number, the garbage collector is run. One can inspect the threshold for new objects (objects in Python known as generation 0 objects) 
        by importing the gc module and asking for garbage collection thresholds.

        - There are two ways for performing manual garbage collection: time-based and event-based garbage collection.
            - Time-based garbage collection is simple: the garbage collector is called after a fixed time interval.
            - Event-based garbage collection calls the garbage collector on event occurrence. For example, when a user exits the application or when the 
            application enters into idle state.

        # ------------- syntax:
        import gc       # loading garbage collection module
        
'''

import gc

print("Garbage collection thresholds: {}".format(gc.get_threshold()))

'''
    # explain output:
        Garbage collection thresholds: (700, 10, 10)
        Here, the default threshold on the above system is 700. This means when the number of allocations vs. the number of deallocations is greater than 700 the 
        automatic garbage collector will run. Thus any portion of your code which frees up large blocks of memory is a good candidate for running manual garbage collection.
'''