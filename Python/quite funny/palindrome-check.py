def palindrome_check(word):
    left, right = 0, len(word) - 1

    while left < right:
        while left < right and not word[left].isalnum():
            left += 1
        while left < right and not word[right].isalnum():
            right -= 1

        if left >= right:
            break

        left_char, right_char = word[left].lower(), word[right].lower()
        if left_char != right_char:
            return f"Mismatch: {word}\nProblem: {left_char} != {right_char}\nIndex of mismatch: ({left} and {right})"

        left += 1
        right -= 1

    return f"Match found: '{word}' is a palindrome"


print(palindrome_check("racecar"))
print(palindrome_check("Hello"))
print(palindrome_check("A man a plan a canal Panama"))
print(palindrome_check("No 'x' in Nixon"))
