#!/usr/bin/python3

'''
    Programs for printing pyramid patterns in Python based on user input
'''

from sys import stdin,stdout

no_of_lines = int(input("Please write your desire number of lines: "))

for line in range(0,no_of_lines):
    for star in range(0,line+1):
        print("*",end="")
    print("\r")
