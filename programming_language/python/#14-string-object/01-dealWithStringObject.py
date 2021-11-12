#!/usr/bin/python3
#coding=utf-8
def main():
    '''
        string.ascii_letters	    Concatenation of the ascii_lowercase and ascii_uppercase constants.
        string.ascii_lowercase	    Concatenation of lowercase letters
        string.ascii_uppercase	    Concatenation of uppercase letters
        string.digits	            Digit in strings
        string.hexdigits	        Hexadigit in strings
        string.lowercase	        A string must contain lowercase letters.

        String.endswith()	        Returns True if a string ends with the given suffix otherwise returns False
        String.startswith()         Returns True if a string starts with the given prefix otherwise returns False
        String.isdigit()	        Returns “True” if all characters in the string are digits, Otherwise, It returns “False”.
        String.isalpha()	        Returns “True” if all characters in the string are alphabets, Otherwise, It returns “False”.
        string.isdecimal()	        Returns true if all characters in a string are decimal.
    '''
    random_text = '''<h2>One advanced diverted domestic sex repeated bringing you old</h2>
<p>Folly words widow one downs few age every seven. If miss part by fact he park just shew. Discovered had get considered projection who favourable. Necessary up knowledge it tolerably. Unwilling departure education is be dashwoods or an. Use off agreeable law unwilling sir deficient curiosity instantly. Easy mind life fact with see has bore ten. Parish any chatty can elinor direct for former. Up as meant widow equal an share least.</p>
<p>Another journey chamber way yet females man. Way extensive and dejection get delivered deficient sincerity gentleman age. Too end instrument possession contrasted motionless. Calling offence six joy feeling. Coming merits and was talent enough far. Sir joy northward sportsmen education. Discovery incommode earnestly no he commanded if. Put still any about manor heard.</p>
<p>Village did removed enjoyed explain nor ham saw calling talking. Securing as informed declared or margaret. Joy horrible moreover man feelings own shy. Request norland neither mistake for yet. Between the for morning assured country believe. On even feet time have an no at. Relation so in confined smallest children unpacked delicate. Why sir end believe uncivil respect. Always get adieus nature day course for common. My little garret repair to desire he esteem.</p>
    '''

    print("[+] str.isdigit():", end=" ")
    print(random_text.isdigit())

if __name__ == "__main__":
    main()