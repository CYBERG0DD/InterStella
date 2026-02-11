def play_palindrome_game():

    score = 0
    trials = 0
    used_words = set() 
    max_trials = 10
    print()
    print("==================== Welcome to the Palindrome Game! ===================")
    ask = input("Would you love to play? (yes/no): ").lower()
    
    if ask != "yes":
        print("Maybe next time! See ya")
        return

    name = input("Enter your username: ")
    while not name.strip():
        name = input("Name is required! Please enter username: ")

    print(f"\nOkay {name}, you have {max_trials} trials. The goal is to find as many palindromes as you can!. LET'S GOOOOO")

    # --- STEP 2: THE GAME LOOP ---
    while trials < max_trials:
        word = input(f"\n[{trials + 1}/{max_trials}] Enter keyword: ").strip().lower()

        # Check for empty input
        if not word:
            print("Don't waste a trial on an empty string! Try again.")
            continue

        if word in used_words:
            print(f"You already used '{word}'! Think of a new one.")
            continue

        trials += 1
        used_words.add(word) 

        cleaned = word.replace(" ", "")
        if cleaned == cleaned[::-1]:
            print(f"✅ Yes! '{word}' is a palindrome.")
            score += 1
        else:
            print(f"❌ No! '{word}' is not a palindrome.")


        if trials == 8:
            print("⚠️ WARNING: Only 2 trials left! use them wisely")
        if trials == 9:
            print("Last trial")

    print("\n" + "="*40)
    print(f"GAME OVER, {name}!")
    print("="*40)
    print("Game Stats: ")
    print()
    print(f"Total Palindromes Found: {score}")
    trials_left = trials
    print(f"Trials left: {trials}")

    if score >= 8:
        print("🏆 Great job! Well done!")
    else:
        print("📉 You did not reach the expected end. Keep practicing!")

play_palindrome_game()