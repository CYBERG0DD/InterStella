# Function Declaration

## Overview
In Go, a function must have a name, parameters with correct 
types, and a return type.

## Function Signatures — Passing Correct Types
```go
// ✅ Correct function signature
func Checkword(text string) string {
    return text
}
// Takes a string and returns a string

// ❌ Wrong — passing wrong type
func Checkword(text int) string {
    return text
}
// text is int but function expects string
```

## Variable Scope
Variables only live inside the block they were declared in.
```go
func greet() {
    name := "MK"  // name only exists inside greet()
}

fmt.Println(name)  // ❌ will not compile
// name is not accessible outside greet()
```

## Passing []string vs string — Type Mismatch
```go
func capitalize(word string) string { ... }

capitalize([]string{"hello", "world"})  // ❌ type mismatch
// Function expects a string but received a []string
```

## string() — Byte to String Conversion
```go
word := "golang"
last := word[len(word)-1]         // byte (not a string)
last  = string(word[len(word)-1]) // ✅ converted to string
```

## Key Rule
> Always pass the correct type a function expects.
> Variables declared inside a function cannot be used outside it.