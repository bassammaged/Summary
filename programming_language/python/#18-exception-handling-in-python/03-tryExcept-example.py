#!/usr/bin/python3

def divison(a,b):
    result  = a/b
    return result


def main():
    numbers = list(map(int,input('Insert 2 number seperated by space: ').split(' ')))
    
    try:
        result  = divison(*numbers)
        print("[+] OUTPUT: Result = {}".format(result))
    except TypeError:
        print("Please insert 2 numbers only.")
    except ZeroDivisionError:
        print("Division by zero, we cannot do that here.")
    # except (ZeroDivisionError,TypeError):
        # print("Error occuared.")
    else:
        print("[+] Notification: Thanks God for no exceptions have been raised.")
    finally:
        print("[+] Close: Thanks and thanks again.")
        
if __name__ == "__main__":
    main()