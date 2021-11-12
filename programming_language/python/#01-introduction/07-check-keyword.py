#!c:\Python390
from keyword import iskeyword
def main():
    keywordsList    = [
        "for",
        "while",
        "tanisha",
        "break"
    ]

    for word in keywordsList:
        if iskeyword(word):
            print("[+]",word,"is python keyword")
        else:
            print("[+]",word, "is not python keyword")
if __name__ == "__main__":
    main()