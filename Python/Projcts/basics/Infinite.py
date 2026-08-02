import sys
import re


def main():
    print("WELCOME TO Berverly Hills.")
    print("[Motto: Comfy homes]")
    print("\nPlease fill in the required information below so we can begin processing your details.\n")

    # --- BASIC INFO ---
    first_name = input("Enter first name: ").strip()
    middle_name = input("Middle name [optional]: ").strip()
    last_name = input("Enter last name: ").strip()

    print(
        f"\nHello, {first_name.title()} {middle_name.title()} {last_name.title()}, and welcome aboard!\n")

    # --- ACCOUNT CREATION ---
    while True:
        username = input("Enter username or [nickname]: ").strip()
        if len(username) < 5:
            print("Invalid. Username must be at least 5 characters long.")
            continue

        email = input("Enter email address: ").strip()
        if "@" not in email or "." not in email:
            print("Invalid email address. Please try again.")
            continue

        # --- PHONE ---
        while True:
            phone = input("Enter phone number (11 digits): ").strip()
            if len(phone) != 11:
                print("Invalid. Must be exactly 11 digits.")
            elif not phone.isdigit():
                print("Invalid. Phone must contain only numbers.")
            else:
                break

        # --- PASSWORD ---
        suggest = input("Suggest strong password? (yes/no): ").strip().lower()
        if suggest == "yes":
            print("Suggested password: I7baS3%&*")
            print("You can save your password to your device or Google for recovery.")
        else:
            print(
                "We recommend using a strong password — your account could be vulnerable otherwise.")
            while True:
                password = input("Create password (8–10 characters): ").strip()
                if len(password) < 8:
                    print("Invalid. Password should be at least 8 characters.")
                elif len(password) > 10:
                    print("Invalid. Password should not exceed 10 characters.")
                else:
                    print("Password accepted.")
                    break

        # --- ID VALIDATION ---
        print("\nYour ID will be sent to your email. Don't share it with anyone.")
        print("Note: It cannot be changed once issued.\n")

        # --- AGE VALIDATION ---
        while True:
            try:
                birth_year = int(input("Enter birth year (YYYY): ").strip())
                age = 2025 - birth_year
                print(f"You are {age} years old.")
                if age < 18 or age > 45:
                    print("Sorry, your age does not meet our requirement (18–45).")
                    continue
                break
            except ValueError:
                print("Please enter a valid year.")

        # --- GENDER ---
        gender = input("Gender (M/F/Other): ").strip().title()

        # --- PROFESSION ---
        professions = [
            "cyber security specialist",
            "software engineer",
            "hardware engineer",
            "it technician",
            "financial management",
            "customer care",
        ]

        print("\nFIELD OF OPERATION:")
        for p in professions:
            print(f"- {p.title()}")

        while True:
            choice = input(
                "\nEnter your profession exactly as listed: ").strip().lower()
            if choice not in [p.lower() for p in professions]:
                print(
                    "Sorry, that profession is not on the list. Please check your spelling.")
                continue
            break

        # --- STRENGTH MARK ---
        while True:
            try:
                mark_str = input(
                    "Rate your field strength (100, 95, 90, 85, 80, 75, 70): ").strip()
                mark = int(mark_str)
                if mark < 70:
                    print("Sorry, that's below our acceptable range.")
                    continue
                break
            except ValueError:
                print("Please enter a valid number.")

        # Strength feedback
        if mark == 100:
            print(
                "You are an expert in this field — we need the best of your knowledge to help us grow.")
        elif mark >= 95:
            print("Excellent knowledge! We trust you'll help in our company's growth.")
        elif mark >= 90:
            print("Strong skills! You'll be a great asset to our company.")
        elif mark >= 85:
            print("Good knowledge — your expertise will help us move forward.")
        elif mark >= 75:
            print("Solid understanding. You'll make our company outstanding.")
        else:
            print("Good — we're happy to have you onboard.")

        # --- STUDENT CHECK ---
        student = input("\nAre you a student (yes/no)?: ").strip().lower()
        if student == "yes":
            print("\nWelcome, young star! You've come to the right place.")
            print("We will begin processing your details and contact you shortly.")
            print("You will receive an email with your job qualification status.")
        else:
            print(
                "\nThank you for your cooperation. We will begin processing your details shortly.")

        print("\n✅ Registration complete! Welcome to GLOBAL '9' CORPORATION.")
        break


if __name__ == "__main__":
    main()
