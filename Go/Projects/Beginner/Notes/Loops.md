# Loops

####  LESSON: Getting characters from a string
####  Learned: March 2026

## Range over a Slice
Iterates over each element in a slice.
````go
words := []string{"hello", "world", "golang"}
for _, word := range words {
    fmt.Println(word)
}
````

## Using _ to Discard Index
`_` tells Go you are declaring a variable but won't use it.
Go throws an error if you declare a variable and never use it.
````go
for _, word := range words { // ✅ index ignored
    fmt.Println(word)
}

for i, word := range words { // ❌ error if i is never used
    fmt.Println(word)
}
````

## Range over a Map
````go
tags := map[string]bool{
    "(up)": true,
    "(low)": true,
}
for key, value := range tags {
    fmt.Println(key, value)
}
````

## Key Rule
> Go will not compile if you declare a variable and never use it.