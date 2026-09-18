# Testing how to write output into a file "redirection" rather that the standard output termial display.
# Using the > for redirection and >> to append to an existing file

def greet_User(name):
    print(f'Welcome {name} CEO of Infinite Technologies')

greet_User('Mel')

# Finding the square root of 9
def find_Square_Root(number):
    total = int(number / 9)
    print(f'The square root of {number} is: {total}')
find_Square_Root(81)

# finding the square using a different approach
def the_square(value):
    final = int(value) * int(value)
    print(f'The square of {value} is: {final}')
the_square(16)

a = ''
print(f'My name in caps: {a.upper()}')

Interger = 16
String = 'Mel CEO OF INFINITE TECHNOLOGIES'
print(type(Interger))
print(type(String))



