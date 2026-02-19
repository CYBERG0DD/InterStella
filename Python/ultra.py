print("Before we begin please enter your name so i would know you better. ")
name = input("\nWhat's your name: ")
print(f"Hello, {name} what would you love to know. ")
choice = input("your (Age) or (birth Year)? ")
if choice == "age":
    year = input("Input birth year: ")
    new_age = 2025 - int(year)
    print(
        f"you are {new_age} years old. You have lived {new_age} years on earth. ")
    puzzle = input('Would you want some friendly tips for your age (yes/no): ')
    if puzzle == 'yes' or 'ok':

        if new_age < 18:
            print("\nyou're still a young growing teenager.  ")
        print('''Frienly "TIPS 4 YOU".
            Be a good and respectful guy/lady
            Stay outta trouble
            No bad friends. ''')

    elif new_age == 18:
        print("you're a matured teenager, with good vibez. ")
        print('''\nFriendly "TIPS 4 YOU" 
                NO (smoking, drinking, clubing, girfriend or boyfriend, gambling and fraud) 
                Obey the law 
                Hustle
                Be responsible
                Stay focused on the dream, the goal and the money . so don't be distracted
                Make God a friend (top priority). ''')

    elif new_age > 18:
        print("At this age you are a already a young adult. ")
        print('''\nFriendly TIPS 4 YOU
                Spend money wisely
                don't chase women
                workout
                be at peace with all men
                No sex or romantic relationship until you're married
                Stay healthy and enjoy life at it's best to the fullest. ''')
    else:
        print('Ok goodbye')

elif choice == "year" or "birthyear":
    age = input("Input age: ")
    new_year = 2025 - int(age)
    print(f"you were born in the year {new_year}. ")
    if new_year < 2000:
        print("you're getting old. ")
    elif new_year >= 2000:
        print("nice, keep up youngstar. ")
question = input('Is that all you would love to know?: ')
if question == 'yes':
    print('oh that\'s great to hear am glad i could help answer some of your questions. Have a nice day!! :) ')
elif question == 'no':
    sef = input('oh, am sorry but that will be all for now: ')
    print('i was only designed to help users with their year of birth and age. ')
