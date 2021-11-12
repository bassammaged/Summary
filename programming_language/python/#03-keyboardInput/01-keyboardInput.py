#!c:\Python390
def main():
    '''
        The input function prompts text if a parameter is given. The functions reads input from the keyboard, converts it to a string and removes the newline.
    '''
    # ----------- taking one input at a time
    keyboardInput   = input("Enter What you want: ")
    print("You instered the following:",keyboardInput,",the type of data:",type(keyboardInput))

    # ----------- taking multiple inputs at a time
    key1,key2,key3  = input('Insert 3 inputs in a row: ').split()       # split by space
    print('taking multiple inputs at a time:',key1,key2,key3)


    # ----------- taking multiple inputs at a time as list
    multipleInputs  = list(input("Enter multiple inputs: ").split())
    print("You instered the following:",multipleInputs,",the type of data:",type(multipleInputs))

    # ----------- taking multiple inputs at a time as list (integer)
    multipleInputs2 = list(map(int,input("Enter multiple integer inputs: ").split()))
    print("You instered the following:",multipleInputs2,",the type of data:",type(multipleInputs2))

if __name__ == "__main__":
    main()