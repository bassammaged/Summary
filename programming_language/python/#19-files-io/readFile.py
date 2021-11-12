#!C:\python39

def main():
    try:
        # ----- Open the file for reading.
        readFile    = open('./from1To1000',mode='r',encoding='utf-8')
        for line in readFile:
            print(line.strip())
            break
    finally:
        # ----- closing the file.
        readFile.close()
if __name__ == '__main__':
    main()