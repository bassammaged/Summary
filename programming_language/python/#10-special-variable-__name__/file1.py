print("file1.py __name__ = {}".format(__name__))

if __name__ == "__main__":
    print("file1.py is being run directly.")
else:
    print("file1.py is being imported.")