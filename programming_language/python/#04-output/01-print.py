def main():
    '''
        - Definition:
        The print() function prints the specified message to the screen, or other standard output device.
        - syntax:
            print(object(s), sep=separator, end=end, file=file, flush=flush)
    '''
    # ----------- print with newline
    print("Hello the end will be newline")

    # ----------- multiple print
    print("Hello","The","Output!")

    # ----------- multiple print with specific separator
    print("Hello","The","Output!",sep="-")

    # ----------- specify the print's end
    print("Hello First Line",end=" ")
    print("Hello Second Line")  

    # ----------- combining positional and keyword arguments 
    print('Number one portal is {}, {}, and {}.'.format('Geeks', 'For','Geeks')) 

    # ----------- combining positional and keyword arguments 
    print('Number one portal is {0}, {1}, and {other}.'.format('Geeks', 'For', other ='Geeks')) 

if __name__ == "__main__":
    main()