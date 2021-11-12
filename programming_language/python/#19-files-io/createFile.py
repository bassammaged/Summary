#!C:\python39

def main():
    
    try:
        # ----- create the file with open function
        createFile  = open('from1To1000',mode='wt',encoding='utf-8')

        for i in range(1,1001):
            createFile.write('Line no.: ' + str(i) + '\n')
    finally:
        # ----- Close the file.
        createFile.close()
if __name__ == "__main__":
    main()