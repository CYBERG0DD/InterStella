import datetime

# Get current date and time dynamically
now = datetime.datetime.now()
date = now.strftime("%d/%m/%Y")
time = now.strftime("%H:%M:%S")

print("\nWELCOME TO THE GLO CAFE APP.")
user_type = input("User: new or existing/old user: ").strip().lower()


# Helper functions
def validate_phone():
    while True:
        phone = input("Enter mobile number (11 digits): ").strip()
        if len(phone) != 11 or not phone.isdigit():
            print("❌ Invalid phone number. Try again.")
        else:
            return phone


def buy_data_menu():
    bundles = {
        "1": ("N200", "500MB"),
        "2": ("N300", "900MB"),
        "3": ("N500", "1GB"),
        "4": ("N1000", "2.5GB"),
        "5": ("N2000", "3GB"),
        "6": ("N4000", "7GB"),
        "7": ("N8000", "12.5GB"),
    }

    print("\nSelect Amount/Data bundle:")
    for key, (price, data) in bundles.items():
        print(f"{key}. {price} for {data}")

    while True:
        select = input(" (1/2/3/4/5/6/7): ").strip()
        if select not in bundles:
            print("Invalid choice, try again.")
            continue

        price, data = bundles[select]
        print(f"\n✅ Recharge of {price} for {data} was successful.")
        print(f"Expires on {date} by {time}.")
        print("Enjoy streaming, browsing, and connecting with loved ones!")
        break


def borrow_data_menu():
    options = {
        "1": ("N50", "45MB"),
        "2": ("N100", "130MB"),
        "3": ("N200", "500MB"),
        "4": ("N300", "900MB"),
        "5": ("N500", "1GB"),
        "6": ("N1000", "2GB"),
    }

    print("\nSelect Data Borrow Option:")
    for key, (price, data) in options.items():
        print(f"{key}. {price} for {data}")

    while True:
        choice = input(" (1/2/3/4/5/6): ").strip()
        if choice not in options:
            print("Invalid option.")
            continue

        price, data = options[choice]
        print(
            f"\n✅ You borrowed {price} for {data}. It expires on {date} by {time} and will be deducted from your next recharge."
        )
        break


def night_life_menu():
    bundles = {
        "1": ("N50", "500MB"),
        "2": ("N100", "700MB"),
        "3": ("N500", "1.5GB"),
    }
    print("\n🌙 GLO NIGHT LIFE BUNDLES (10:59 PM - 6:00 AM)")
    for key, (price, data) in bundles.items():
        print(f"{key}. {price} for {data}")

    while True:
        choice = input(" (1/2/3): ").strip()
        if choice not in bundles:
            print("Invalid choice.")
            continue

        price, data = bundles[choice]
        print(f"\n✅ {price} for {data} activated! Expires by 6:00 AM.")
        break


# ---- Main Flow ----
if user_type == "new":
    mobile = validate_phone()
    print("An OTP has been sent to your phone.")
    otp = input("Enter OTP: ")
    password = input("Create a 4-digit password: ")
    print(
        f"\n🎉 Welcome to the Glo Café App, {mobile}! You've received 75GB of free data valid until 12AM!"
    )

elif user_type in ["old", "existing"]:
    username = input("Username: ")
    password = input("Password: ")
    print(f"\nWelcome back, {username}!")
    print("What would you like to do today?")
    print(
        "1. Buy Airtime\n2. Buy Data\n3. Buy for a Friend\n4. Borrow Airtime/Data\n5. Glo Night Life"
    )

    choice = input("Choose (1/2/3/4/5): ").strip()

    if choice == "1":
        amount = input("Enter amount: ₦")
        print(f"✅ Recharge of ₦{amount} was successful!")

    elif choice == "2":
        buy_data_menu()

    elif choice == "3":
        recipient = validate_phone()
        amount = input("Enter amount to send: ₦")
        print(
            f"✅ Transfer of ₦{amount} to {recipient} was successful! The recipient will receive it shortly."
        )

    elif choice == "4":
        print(
            "\nWelcome to *Glo Borrow Me* service. You can also dial *399# directly."
        )
        print("1. Borrow Airtime\n2. Borrow Data")
        option = input(" (1/2): ").strip()
        if option == "1":
            amt = input("Enter amount to borrow: ₦")
            print(
                f"✅ You borrowed ₦{amt}. It will be deducted from your next recharge."
            )
        elif option == "2":
            borrow_data_menu()

    elif choice == "5":
        night_life_menu()

    else:
        print("Invalid option. Goodbye!")

else:
    print("❌ Invalid input. Please restart and select either 'new' or 'old' user.")
