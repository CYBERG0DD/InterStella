# Dealing With Slice

####  LESSON: Getting characters from a string
####  Learned: March 2026

```Go
package main


func main() {
    word := "golang"

    // first character
    _ = word[:1]  // "g"

    // last character (must convert byte → string)
    _ = string(word[len(word)-1])  // "g"

    // everything except last character
    _ = word[:len(word)-1]  // "golan"

    // middle (excluding first and last)
    _ = word[1:len(word)-1]  // "olan"
}
```

## Getting characters from a string (Tabular Version)

| Slice | Result |
|---|---|
| `word[:1]` | First character |
| `word[1:]` | Everything after first |
| `word[:len(word)-1]` | Everything before last |
| `word[len(word)-1]` | Last character (byte) |
| `string(word[len(word)-1])` | Last character (string) |

## Example
```go
word := "golang"
first  := word[:1]                         // "g"
last   := string(word[len(word)-1])        // "g"
middle := word[1:len(word)-1]              // "olan"
```