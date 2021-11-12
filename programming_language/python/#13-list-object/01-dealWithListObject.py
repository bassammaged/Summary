#!/usr/bin/python3
#coding=utf-8
def main():
    '''
    # ----------- Structure List Methods
    append()	Add an element to the end of the list
    extend()	Add all elements of a list to the another list
    insert()	Insert an item at the defined index
    remove()	Removes an item from the list
    pop()	    Removes and returns an element at the given index
    clear()	    Removes all items from the list
    count()	    Returns the count of number of items passed as an argument
    sort()	    Sort items in a list in ascending order
    reverse()	Reverse the order of items in the list
    copy()	    Returns a copy of the list
    
    # ----------- Search List Methods
    index()	    Returns the index of the first matched item

    # ----------- Casting List Methods
    ', '.join(listVar)                                                    # return string instead of list 
    stringVar.split(' - ')                                                # return list instead of string


    # ----------- Math Built-in functions with List
    sum()	    Sums up the numbers in the list
    ord()	    Returns an integer representing the Unicode code point of the given Unicode character
    cmp()	    This function returns 1, if first list is “greater” than second list
    max()	    return maximum element of given list
    min()	    return minimum element of given list
    

    # ----------- Dealing Built-in functions with List
    len()	    Returns length of the list or size of the list
    enumerate()	Returns enumerate object of list
    map()	    Returns a list of the results after applying the given function to each item of a given iterable
    '''

    random_list = [1,10,3,5.4]
    
    

    print("[+] list.index():", end=" ")
    print(random_list.index(10))

if __name__ == "__main__":
    main()