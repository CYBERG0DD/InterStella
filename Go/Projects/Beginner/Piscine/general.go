package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func hex2dec(hexStr string) (int64, error) {
	return strconv.ParseInt(hexStr, 16, 64)
}

func bin2dec(binStr string) (int64, error) {
	return strconv.ParseInt(binStr, 2, 64)
}

func capitalizefirst(sentence string) string {
	return strings.ToUpper(sentence[:1]) + strings.ToLower(sentence[1:])
}

func capitalizetwo(set []string, n int) []string {
	for char := range set {
		if char >= len(set)-n {
			set[char] = strings.ToUpper(set[char])
		}
	}
	return set
}

func fixpunctuation(strin []string) string {
	result := strings.Join(strin, "")
	return strings.ReplaceAll(result, ",", ", ")
}

func ispuncutation(s string) bool {
	for _, i := range s {
		if unicode.IsPunct(i) {
			return true
		}
	}
	return false
}

func aORan(word string) string {
	if strings.ContainsRune("aeiouHAEIOUH", rune(word[0])) {
		return "an"
	}
	return "a"
}

func fixarticles(word string) string {
	re := regexp.MustCompile(`'\s*(.*?)\s*'`)
	return re.ReplaceAllString(word, "'$1'")
}

func a2an(text string) string {
	text = strings.ReplaceAll(text, "A", "An")
	text = strings.ReplaceAll(text, "An book", "A book")
	return text
}

func main() {
	val, _ := hex2dec("1E")
	val1, _ := bin2dec("10")

	fmt.Println(val, val1)

	words := []string{"hello", "world", "from", "golang"}
	for _, word := range words {
		fmt.Println(capitalizefirst(word))
	}

	fmt.Println(capitalizetwo([]string{"This", "is", "so", "exciting"}, 2))
	fmt.Println(fixpunctuation([]string{"Hello", ",", "World!"}))

	fmt.Println(ispuncutation("!"))
	fmt.Println(ispuncutation("."))
	fmt.Println(ispuncutation("y"))

	fmt.Println(aORan("apple"))
	fmt.Println(aORan("horse"))
	fmt.Println(aORan("An man, was eating a apple"))

	text := "' awesome '"
	word := "' Hello, World! '"
	fmt.Println(fixarticles(text))
	fmt.Println(fixarticles(word))

	fmt.Println(a2an("There it was. A amazing rock. A honest man. A book"))
}
