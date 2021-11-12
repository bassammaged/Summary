
#!/usr/bin/python3

'''
    # ---------- Where they are used?
        Different looping techniques are primarily useful in the places where we don’t need to actually manipulate the structure and order 
        of the overall containers, rather only print the elements for a single-use instance, no in-place change occurs in the container. This can 
        also be used in instances to save time.

    # ---------- Using enumerate():
        enumerate() is used to loop through the containers printing the index number along with the value present in that particular index.

    # ---------- Using zip():
        zip() is used to combine 2 similar containers(list-list or dict-dict) printing the values sequentially. The loop exists only till the 
        smaller container ends.

    
    # ---------- Using iteritem(): Deprecated in python 3.x
        iteritems() is used to loop through the dictionary printing the dictionary key-value pair sequentially. 
    
    # ---------- Using items():
        items() performs the similar task on dictionary as iteritems() but have certain disadvantages when compared with iteritems().
        It is very time-consuming. Calling it on large dictionaries consumes quite a lot of time.   (in python 2.x)
        It takes a lot of memory. Sometimes takes double the memory when called on a dictionary.    (in python 2.x)
'''

# enumerate()
titles    = ['Pen Tester','Info Sec Eng','Cyber Sec Eng','SOC Analyst','Network Sec Eng']
print("[+] enumerate() function:")
for key,value in enumerate(titles):
    print('{}:{}'.format(key,value))

# zip()
departments  = ["Information security","Cyber Security","Security Operation"]
print("[+] zip() function:")
for value1,value2 in zip(titles,departments):
    print('{}:{}'.format(value1,value2))

# iteritems() & items()
employees   = {100:"Moscu",101:"CPUKiller",102:"Hamada"}

print("[+] iteritems() & items() function:")
for key,value in employees.items():
    print ("{}:{}".format(key,value))


