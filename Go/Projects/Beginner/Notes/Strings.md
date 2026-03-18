# Strings

####  LESSON: Getting characters from a string
####  Learned: March 2026


## Overview
The strings package provides functions for string manipulation in Go.

## Functions

### strings.ToUpper()
Converts a string to ALL CAPS.
```go
strings.ToUpper("hello")  // "HELLO"
```

### strings.ToLower()
Converts a string to all lowercase.
```go
strings.ToLower("HELLO")  // "hello"
```

### strings.Fields()
Splits a string by whitespace into a slice.
```go
strings.Fields("hello world")  // ["hello", "world"]
```

### strings.Split()
Splits a string by a specific separator.
```go
strings.Split("a,b,c", ",")  // ["a", "b", "c"]
```

### strings.Join()
Joins a slice of strings with a separator.
```go
strings.Join([]string{"hello", "world"}, " ")  // "hello world"
```

### strings.ContainsAny()
Checks if a string contains any of the given characters.
```go
strings.ContainsAny("hello", "aeiou")  // true
```

### strings.TrimSpace()
Removes leading and trailing spaces.
```go
strings.TrimSpace("  hello  ")  // "hello"
```