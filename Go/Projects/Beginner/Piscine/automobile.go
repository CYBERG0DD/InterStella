package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// --- DATA STRUCTURES ---

type Bike struct {
	Name         string
	Manufacturer string
	Country      string
	Category     string // Sport, Adventure, Tour
	EngineCC     int
	SkillLevel   string // Beginner, Intermediate, Expert
}

// --- THE GIGANTIC DATABASE ---
var MasterDatabase = []Bike{
	// JAPAN
	{Name: "CBR500R", Manufacturer: "Honda", Country: "Japan", Category: "Sport", EngineCC: 471, SkillLevel: "Beginner"},
	{Name: "CBR1000RR-R", Manufacturer: "Honda", Country: "Japan", Category: "Sport", EngineCC: 999, SkillLevel: "Expert"},
	{Name: "Ninja 400", Manufacturer: "Kawasaki", Country: "Japan", Category: "Sport", EngineCC: 399, SkillLevel: "Beginner"},
	{Name: "Ninja ZX-10R", Manufacturer: "Kawasaki", Country: "Japan", Category: "Sport", EngineCC: 998, SkillLevel: "Expert"},
	{Name: "Africa Twin", Manufacturer: "Honda", Country: "Japan", Category: "Adventure", EngineCC: 1084, SkillLevel: "Intermediate"},
	{Name: "Versys 650", Manufacturer: "Kawasaki", Country: "Japan", Category: "Tour", EngineCC: 649, SkillLevel: "Intermediate"},

	// INDIA
	{Name: "Himalayan 450", Manufacturer: "Royal Enfield", Country: "India", Category: "Adventure", EngineCC: 452, SkillLevel: "Beginner"},
	{Name: "Interceptor 650", Manufacturer: "Royal Enfield", Country: "India", Category: "Tour", EngineCC: 648, SkillLevel: "Intermediate"},
	{Name: "RC 390", Manufacturer: "KTM", Country: "India", Category: "Sport", EngineCC: 373, SkillLevel: "Intermediate"},
	{Name: "Apache RR 310", Manufacturer: "TVS", Country: "India", Category: "Sport", EngineCC: 312, SkillLevel: "Beginner"},

	// GERMANY
	{Name: "S1000RR", Manufacturer: "BMW", Country: "Germany", Category: "Sport", EngineCC: 999, SkillLevel: "Expert"},
	{Name: "R 1300 GS", Manufacturer: "BMW", Country: "Germany", Category: "Adventure", EngineCC: 1300, SkillLevel: "Expert"},
	{Name: "K 1600 GTL", Manufacturer: "BMW", Country: "Germany", Category: "Tour", EngineCC: 1649, SkillLevel: "Expert"},

	// UNITED KINGDOM (UK)
	{Name: "Tiger 900", Manufacturer: "Triumph", Country: "UK", Category: "Adventure", EngineCC: 888, SkillLevel: "Intermediate"},
	{Name: "Speed Triple 1200", Manufacturer: "Triumph", Country: "UK", Category: "Sport", EngineCC: 1160, SkillLevel: "Expert"},
	{Name: "Rocket 3", Manufacturer: "Triumph", Country: "UK", Category: "Tour", EngineCC: 2458, SkillLevel: "Expert"},

	// UNITED STATES (USA)
	{Name: "Pan America 1250", Manufacturer: "Harley-Davidson", Country: "USA", Category: "Adventure", EngineCC: 1252, SkillLevel: "Expert"},
	{Name: "Road Glide", Manufacturer: "Harley-Davidson", Country: "USA", Category: "Tour", EngineCC: 1868, SkillLevel: "Intermediate"},
	{Name: "FTR 1200", Manufacturer: "Indian", Country: "USA", Category: "Sport", EngineCC: 1203, SkillLevel: "Expert"},

	// CHINA
	{Name: "450SR", Manufacturer: "CFMOTO", Country: "China", Category: "Sport", EngineCC: 450, SkillLevel: "Beginner"},
	{Name: "800MT", Manufacturer: "CFMOTO", Country: "China", Category: "Adventure", EngineCC: 799, SkillLevel: "Intermediate"},

	// AUSTRIA
	{Name: "890 Adventure R", Manufacturer: "KTM", Country: "Austria", Category: "Adventure", EngineCC: 889, SkillLevel: "Expert"},
	{Name: "1290 Super Duke R", Manufacturer: "KTM", Country: "Austria", Category: "Sport", EngineCC: 1301, SkillLevel: "Expert"},
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// --- 1. USER REGISTRATION ---
	fmt.Println("\nHello and welcome all Motorcycle lovers and enthusiasts!")
	fmt.Println("PLEASE FILL IN THE FOLLOWING INFO...")
	fmt.Println("*****************************************")

	var fullName, email string
	var age int

	// (Name Validation)
	for {
		fmt.Print("Enter full name: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Print("Error reading input", err)
		}
		fullName = strings.TrimSpace(input)
		if fullName == "" {
			fmt.Print("Name is required!\n")
			continue
		}
		isValid := true
		for _, item := range fullName {
			if !unicode.IsLetter(item) && !unicode.IsSpace(item) {
				fmt.Println("Name must only contain letters and spaces!")
				isValid = false
				break
			}
		}

		if isValid {
			break
		}
	}

	// (Email Validation)
	for {
		fmt.Print("Enter email: ")
		input, _ := reader.ReadString('\n')
		email = strings.TrimSpace(input)
		if !strings.Contains(email, "@") {
			fmt.Println("Invalid email!")
			continue
		}
		break
	}

	// (Age Validation)
	for {
		fmt.Print("Enter age: ")
		input, _ := reader.ReadString('\n')
		val, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil || val < 16 || val > 100 {
			fmt.Println("Please enter a valid age (16-100)!")
			continue
		}
		age = val
		break
	}

	// --- 2. PREFERENCE SELECTION ---

	// Country Selection
	countries := []string{"Japan", "India", "Germany", "UK", "USA", "China", "Austria"}
	fmt.Println("\nSelect Country of Interest:")
	for i, c := range countries {
		fmt.Printf("%d. %s\n", i+1, c)
	}

	var selectedCountry string
	for {
		fmt.Print("Choice (1-7): ")
		input, _ := reader.ReadString('\n')
		idx, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil || idx < 1 || idx > len(countries) {
			fmt.Print("Invalid selection!\n")
			continue
		}
		selectedCountry = countries[idx-1]
		break
	}

	// Style Selection
	styles := []string{"Sport", "Adventure", "Tour"}
	fmt.Println("\nWhat is your riding style?")
	for i, s := range styles {
		fmt.Printf("%d. %s\n", i+1, s)
	}

	var selectedStyle string
	for {
		fmt.Print("Choice (1-3): ")
		input, _ := reader.ReadString('\n')
		idx, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil || idx < 1 || idx > 3 {
			fmt.Print("Invalid selection!\n")
			continue
		}
		selectedStyle = styles[idx-1]
		break
	}

	// --- 3. THE RECOMMENDATION ENGINE ---
	fmt.Printf("\n--- RECOMMENDATIONS FOR %s ---\n", strings.ToUpper(fullName))
	fmt.Printf("Criteria: %s bikes from %s\n", selectedStyle, selectedCountry)
	fmt.Println("*****************************************")

	found := false
	for _, bike := range MasterDatabase {
		if bike.Country == selectedCountry && bike.Category == selectedStyle {
			// Skill recommendation logic based on user age and bike CC
			skillAdvice := ""
			if age < 21 && bike.EngineCC > 600 {
				skillAdvice = "[WARNING: High power for your age group!]"
			} else if bike.SkillLevel == "Expert" {
				skillAdvice = "[Requires advanced track training]"
			} else {
				skillAdvice = "[Great for your profile]"
			}

			fmt.Printf("Model: %-20s | CC: %-5d | Skill: %-10s %s\n",
				bike.Name, bike.EngineCC, bike.SkillLevel, skillAdvice)
			found = true
		}
	}

	if !found {
		fmt.Println("No exact matches found for this combination in our database.")
	}

	fmt.Println("\nThank you for using the Moto-Selector! Safe riding.")
}
