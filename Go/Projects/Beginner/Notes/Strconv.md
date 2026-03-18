# String Conversion

## Overview
The strconv package handles conversions between strings and numbers in Go.

## Functions

### strconv.ParseInt()
Converts a string to int64 given a base and bit size.
```go
result, err := strconv.ParseInt("1E", 16, 64)  // 30
result, err := strconv.ParseInt("10", 2, 64)   // 2
result, err := strconv.ParseInt("17", 8, 64)   // 15
```

### strconv.Itoa()
Converts an integer to a string (int to ASCII).
```go
strconv.Itoa(30)  // "30"
```

### strconv.Atoi()
Converts a string to an integer (ASCII to int).
```go
strconv.Atoi("30")  // 30
```

## Base Reference

| Base | Number System | Example Input | Result |
|------|--------------|---------------|--------|
| 2    | Binary       | "10"          | 2      |
| 8    | Octal        | "17"          | 15     |
| 16   | Hexadecimal  | "1E"          | 30     |

> Any base outside 2, 8 and 16 will throw an error.

## Common Bug
Using base 8 instead of base 16 for hex conversion.
```go
// ❌ Wrong — base 8 (octal) used for hex
strconv.ParseInt("1E", 8, 64)   // error

// ✅ Correct — base 16 for hex
strconv.ParseInt("1E", 16, 64)  // 30
```

## Key Rule
> Always double check your base value when using ParseInt.
> Base 16 for hex, base 2 for binary, base 8 for octal.