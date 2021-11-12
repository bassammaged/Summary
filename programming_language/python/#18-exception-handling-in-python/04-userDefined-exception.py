#!/usr/bin/python3

'''
    # -------- Standard exception:
        Python throws errors and exceptions, when there is a code gone wrong, which may cause program to stop abruptly. Python also provides exception handling 
        method with the help of try-except. Some of the standard exceptions which are most frequent include IndexError, ImportError, IOError, ZeroDivisionError, 
        TypeError and FileNotFoundError. A user can create his own error using exception class.

    # -------- User-defined exception:
        - Programmers may name their own exceptions by creating a new exception class. Exceptions need to be derived from the Exception class, either directly 
        or indirectly. Although not mandatory, most of the exceptions are named as names that end in “Error” similar to naming of the standard exceptions in python.

        - Users can define custom exceptions by creating a new class. This exception class has to be derived, either directly or indirectly, from the built-in Exception class.
    
    # -------- syntax:
        class class_name(exception):
            statement(s)

'''

# Quick Example: Guess the right value

import random

class Error(Exception):
    # Baw class for other exception, will be used for inherent
    pass

class ValueLowerError(Error):
    '''Raised when the input value is lower than expected value.'''
    pass

class ValueGreaterError(Error):
    '''Raised when the input value is graeter than expected value.'''
    pass

class GameLevelError(Error):
    pass

def GenerateGameValue(level):
    '''1:'Easy',2:'Medium',~3:'Hard'''

    if level == 1:
        RangeNumbers    = range(1,5)
    elif level == 2:
        RangeNumbers    = range(1,15)
    elif level == 3:
        RangeNumbers    = range(1,45)
    else:
        raise GameLevelError
    
    GameValue       = random.choice(RangeNumbers)
    return GameValue

def TakeGameDiffiucltyInput():
    PlayerInput = int(input("[+] INPUT: Choose game difficulty 1 for Easy, 2 for Medium, 3 for Hard: "))
    return PlayerInput

def TakeGuessNumberInput():
    PlayerInput = int(input("[+] INPUT: Guess the number: "))
    return PlayerInput

def main():
    while True:
        try:
            print("----------------- Guess IT! -----------------")
            # Determine Game Difficulty level
            level       = TakeGameDiffiucltyInput()
            GameValue   = GenerateGameValue(level)
            
            # Take Guess number
            GuessNumber = TakeGuessNumberInput()

            if GameValue == GuessNumber:
                print("[+] Result: Number you guessed '{}' is equal to the game value '{}'.".format(GuessNumber,GameValue))
                break
            elif GuessNumber < GameValue:
                raise ValueLowerError
            else:
                raise ValueGreaterError
        except GameLevelError:
            print("[+] ERROR: Game difficulty value is wrong, try again!")
        except ValueLowerError:
            print("[+] Result: Number you guessed '{}' is lower than the game value '{}'.".format(GuessNumber,GameValue))
        except ValueGreaterError:
            print("[+] Result: Number you guessed '{}' is greater than the game value '{}'.".format(GuessNumber,GameValue))
        except KeyboardInterrupt:
            print("\n[+] ERROR: Something wrong!")
            break
        else:
            print("[+] Congratulation! mememememem xD.")
        finally:
            print("[+] Notice: We are happy for your trust.")

if __name__ == "__main__":
    main()
