# ----- Import the module
import argparse

# ----- take input from the user:
def describe_argparse():
    # ------ Craete a parser
    parser = argparse.ArgumentParser(description="The main Parser")
        # -- Arguments:
        # description: describe the parser
        # epilog: more description at the end of parser
        # add_help: by default is true to show the pasrer help.
    # ----- Adding the arguments
    parser.add_argument("-d","--domain",help="Help from foo!")
        # -- Arguments
        # name of the flag.
        # help: A brief description of what the argument does.
        # default: The value produced if the argument is absent from the command line and if it is absent from the namespace object.
        # type: The type to which the command-line argument should be converted
        # required: Whether or not the command-line option may be omitted (optionals only).
        # choices: A container of the allowable values for the argument.
        # metavar: show an example of the vaule
        # action: The basic type of action to be taken when this argument is encountered at the command line.

    # ----- Parse the arguments
    args = parser.parse_args()


def real_scenario():
    # Create the parser
    parser = argparse.ArgumentParser()
    
    # remove the optional group
    optional = parser._action_groups.pop()
    # adding required groups
    required = parser.add_argument_group("Required arguments")

    required.add_argument("-d",help="The domain of the target",required=True,metavar="example.com")
    optional.add_argument("-t",help="Multi-thread support, by default the number of thread is 1", metavar="",default=1)

    # adding predefined optional after the required.
    parser._action_groups.append(optional) # added this line

    # parsing the arguments
    parser.parse_args()

def main():
    real_scenario()

if __name__ == "__main__":
    main()