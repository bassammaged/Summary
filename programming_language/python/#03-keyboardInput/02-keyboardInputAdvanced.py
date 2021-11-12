#!C:\Python390

from sys import stdin,stdout        # import sys to speedup input process


# ----------------------- start stdin for int function ----------------------- #
def GetInt():
    return list(map(int,stdin.readline().strip().split()))
# -----------------------   end stdin for int function ----------------------- #

# ----------------------- start stdin for string function ----------------------- #
def GetStr():
    return map(str,stdin.readline().strip().split())
# -----------------------   end stdin for string function ----------------------- #

# ----------------------- start stdin for array function ----------------------- #
def GetList():
    return list(map(int,stdin.readline().strip().split()))
# -----------------------   end stdin for array function ----------------------- #

def main():
    print("Please Insert your grades spreated by space:",end=" ")
    GetIntResult    = GetInt()
    print("stdin value:",GetIntResult)

if __name__ == '__main__':
    main()