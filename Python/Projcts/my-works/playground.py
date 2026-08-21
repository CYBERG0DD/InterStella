Account_balance = 100000

user_name = input('Enter name: ')
print(f'Welcome {user_name} to bike shop where you can select your favorite sport bike\n you have a maximum account balance of {Account_balance}')

sports_bike = ['kawasaki ninja h2r -- $60,000', 'Ducati superleggera -- $100,000', 'Aprila RSV4 1100 factory -- $35,000', 'BMW S1000RR -- $33,000']
print('This are the list of bikes we currently have\n')
for lists, bikes in enumerate(sports_bike, start=1):
    
    print(f'{lists} . {bikes}')

print('Select the bike you want to purchase from the list\n')

while True:
    try:
        choice = int(input('Select bike (numbers only): '))
        break
    except ValueError:
        print('Only numbers are allowed for selecting an option e.g (1,2)')

