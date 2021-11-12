#!C:\Python27
# coding= utf-8
from random import randint

# ----------- Function name as parameter
sec_number   = False
def get_secret_number():
    global sec_number
    sec_number = randint(1,500)
    return sec_number

def main():
    '''
        Vulnerability in input() function – Python 2.x
            - There are two common methods to receive input in Python 2.x:
                Using the input() function: This function takes the value and type of the input you enter as it is without modifying any type.
                Using the raw_input() function : This function explicitly converts the input you give to type string,
            
            - The vulnerability in input() method lies in the fact that the variable accessing the value of input can be accessed by anyone just by using the name of 
              variable or method. Let’s explore these one by one: 
                1- Variable name as input parameter: The variable having the value of input variable is able to access the value of the input variable directly.
                2- Function name as parameter: The vulnerability lies here as we can even provide the name of a function as input and access values that are 
                   otherwise not meant to be accessed.
    '''

    # ----------- Variable name as input parameter
    print("============ Variable name as input parameter ============")
    secret_number   = randint(1,100)
    guess_number    = input("Enter what you guess: ")

    if guess_number == secret_number:
        print("[+] Winning: You won!")
    else:
        print('[+] Losing: try again!')


    # ----------- Function name as parameter (cont.)
    print("============ Function name as parameter ============")
    guess_number    = input("Enter what you guess: ")
    if (sec_number == False):
        print("The game seems to be down! try again later.")
    else:
        if guess_number == sec_number:
            print("[+] Winning: You won!")
        else:
         print('[+] Losing: try again!')



if __name__ == "__main__":
    main()