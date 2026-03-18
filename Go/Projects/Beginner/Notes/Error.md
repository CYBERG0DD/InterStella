# Error Handling

## Overview
In Go, functions that can fail return two values — the result 
and an error. Always check and return the error properly.
```go
func convert(num string) (int64, error) {
    result, err := strconv.ParseInt(num, 16, 64)
    if err != nil {
        return 0, err  // ✅ always return err
    }
    return result, nil
}
```

## Silent Error Swallowing Bug
Returning `nil` instead of `err` hides the error from the caller.
```go
// ❌ Wrong — error is swallowed
if err != nil {
    return 0, nil  // caller thinks everything is fine
}

// ✅ Correct — error is passed up
if err != nil {
    return 0, err  // caller knows something went wrong
}
```

## Key Rule
> Never return nil when an error exists. Always return err 
> so the caller knows something went wrong.