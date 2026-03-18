# Regular Expressions (regexp)

## Overview
The regexp package is used to find and manipulate patterns in strings.
It is useful for cleaning, validating and transforming text.

## regexp.Compile() vs regexp.MustCompile()

| Function | Behaviour |
|---|---|
| regexp.Compile() | Returns (regex, error) — you handle the error |
| regexp.MustCompile() | Panics if pattern is invalid — use when certain |
```go
// regexp.Compile — safe, handle error yourself
re, err := regexp.Compile(`'\s*(.*?)\s*'`)
if err != nil {
    log.Fatal(err)
}

// regexp.MustCompile — panics if pattern is wrong
re := regexp.MustCompile(`'\s*(.*?)\s*'`)
```

## Pattern Breakdown

| Pattern | Meaning |
|---|---|
| `'` | Literal single quote |
| `\s*` | Zero or more spaces |
| `.*?` | Capture anything (non-greedy) |
| `()` | Capture group |
| `$1` | Reference to first capture group |

## ReplaceAllString() — Find and Replace
Replaces every pattern match with the replacement string.
```go
re := regexp.MustCompile(`'\s*(.*?)\s*'`)
result := re.ReplaceAllString("' hello '", "'$1'")
// Output: 'hello'
```

## $1 — Referencing Capture Groups
`$1` refers to whatever was captured inside the first `()` group.
```go
// Pattern:     '\s*(.*?)\s*'
// Input:       ' hello '
// $1 captures: hello
// Output:      'hello'
```

## Key Rule
> Use regexp.MustCompile() when you are certain the pattern is correct.
> Use regexp.Compile() when you want to handle errors gracefully.