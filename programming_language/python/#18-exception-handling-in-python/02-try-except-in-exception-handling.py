#!/usr/bin/python3

'''
    # ----------- What is try and except block?
        The try and except block in Python is used to catch and handle exceptions. Python executes code following the try statement as a “normal” part 
        of the program. The code that follows the except statement is the program’s response to any exceptions in the preceding try clause.

    # ----------- syntax
        try:
            statement(s)
        except (exception):
            statement(s), Handling of exception (if required)
        except:
            statement(s), Handling of exception (if required)
        else:
            statement(s), execute if no exception
        finally:
            statement(s), (always executed)
'''

nums    = [1,2,3]

try:
    print("Fisrt element: {}".format(nums[0]))
    print("fourth element: {}".format(nums[3]))
except IndexError:
# except:           in case of general exception
    print("An error occured. because of the inserted index is not found.")
