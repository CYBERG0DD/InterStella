while True:
    try:
        weight = int(input("Enter your weight: "))
        unit = input("(L)bs or (k)g: ")
        if unit.upper == "L":
            converted = weight * 0.45
            print(f"you're {converted} pounds. ")
        else:
            converted = weight // 0.45
            print(f"you are {converted} kilos. ")
    except ValueError:
        print("invalid input. please enter a number. ")
    else:
        break
