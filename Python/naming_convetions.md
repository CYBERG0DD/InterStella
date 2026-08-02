# Variable Naming Conventions

Variable naming conventions improve code readability, maintainability, and consistency across projects. Different programming languages and frameworks follow different conventions.

---

## Naming Convention Reference

| Convention | Example | Common Use |
|------------|---------|------------|
| **camelCase** | `userName` | Variables, functions, methods (JavaScript, Java, Go local variables) |
| **PascalCase** | `UserName` | Structs, classes, interfaces, types (Go exported identifiers, C#, Java) |
| **snake_case** | `user_name` | Variables and functions in Python, C, Rust |
| **SCREAMING_SNAKE_CASE** | `MAX_BUFFER_SIZE` | Constants, macros, environment variables |
| **kebab-case** | `user-name` | URLs, CSS classes, folder names, file names |
| **Train-Case** | `User-Name` | Titles, documentation, headings |
| **flatcase** | `username` | Package names, domains, Linux commands |
| **UPPERCASE** | `USERNAME` | Environment variables, legacy constants |
| **Hungarian Notation** | `strName`, `iCount`, `bIsActive` | Legacy Windows/C/C++ code |
| **Prefix Notation** | `_userName`, `m_name` | Private members in some languages/frameworks |
| **Suffix Notation** | `countInt`, `nameStr` | Rarely used today |

---

# Examples

## 1. camelCase

### Definition
The first word starts with a lowercase letter. Every subsequent word starts with an uppercase letter.

### Syntax

```text
userName
firstName
totalPrice
isLoggedIn
```

### Example

```go
var userName string
var totalPrice float64

func calculateTax() {}
```

### Common Languages

- Go (local variables)
- JavaScript
- Java
- TypeScript
- Kotlin

---

## 2. PascalCase

### Definition
Every word starts with an uppercase letter.

### Syntax

```text
UserName
CustomerAccount
TotalPrice
```

### Example

```go
type User struct {}

func CalculateTax() {}
```

### Common Languages

- Go (exported identifiers)
- C#
- Java
- TypeScript

---

## 3. snake_case

### Definition
All words are lowercase and separated using underscores.

### Syntax

```text
user_name
first_name
total_price
```

### Example

```python
user_name = "MK"
total_price = 100
```

### Common Languages

- Python
- C
- Rust
- Ruby

---

## 4. SCREAMING_SNAKE_CASE

### Definition
Same as snake_case but every letter is uppercase.

### Syntax

```text
MAX_BUFFER_SIZE
DATABASE_URL
PI
```

### Example

```python
MAX_USERS = 100
DATABASE_URL = "localhost"
```

### Common Uses

- Constants
- Environment Variables
- C Macros

---

## 5. kebab-case

### Definition
Words are separated using hyphens (`-`).

### Syntax

```text
user-name
login-page
my-awesome-project
```

### Common Uses

- URLs
- CSS Classes
- HTML IDs
- Folder Names
- File Names

> **Note:** `kebab-case` cannot be used for variable names in most programming languages because `-` is interpreted as the subtraction operator.

---

## 6. Train-Case

### Definition
Every word starts with an uppercase letter and is separated using hyphens.

### Syntax

```text
User-Profile
Dark-Theme
```

### Common Uses

- Documentation
- Titles
- Headings

---

## 7. flatcase

### Definition
Everything is lowercase without separators.

### Syntax

```text
username
projectname
calculator
```

### Common Uses

- Go Packages
- Linux Commands
- Domain Names

### Example

```go
package calculator
```

---

## 8. UPPERCASE

### Definition
Everything is uppercase without spaces or separators.

### Syntax

```text
USERNAME
PASSWORD
TOKEN
```

### Common Uses

- Environment Variables
- Legacy Constants

---

## 9. Hungarian Notation

### Definition
A prefix indicates the variable's data type.

### Syntax

```text
strName
iCount
bIsActive
fPrice
```

### Prefix Reference

| Prefix | Meaning |
|---------|----------|
| `str` | String |
| `i` | Integer |
| `b` | Boolean |
| `f` | Float |
| `ch` | Character |

### Note

Hungarian notation was popular in older Windows and C/C++ applications but is generally discouraged in modern programming because IDEs already know variable types.

---

## 10. Prefix Notation

### Definition
A special character or letter is placed before the variable name.

### Syntax

```text
_userName
m_name
s_count
```

### Common Uses

- Private members
- Class member variables
- Static variables

---

## 11. Suffix Notation

### Definition
The data type is placed at the end of the variable name.

### Syntax

```text
countInt
nameStr
priceFloat
```

### Common Uses

Rare in modern programming.

---

# Naming Convention Comparison

| Convention | Example | Variables | Functions | Classes/Structs | Constants |
|------------|---------|-----------|-----------|-----------------|-----------|
| camelCase | `userName` | ✅ | ✅ | ❌ | ❌ |
| PascalCase | `UserName` | Sometimes | Sometimes | ✅ | Sometimes |
| snake_case | `user_name` | ✅ | ✅ | ❌ | ❌ |
| SCREAMING_SNAKE_CASE | `MAX_USERS` | ❌ | ❌ | ❌ | ✅ |
| kebab-case | `user-name` | ❌ | ❌ | ❌ | ❌ |
| Train-Case | `User-Name` | ❌ | ❌ | ❌ | ❌ |
| flatcase | `username` | Sometimes | Sometimes | ❌ | ❌ |
| UPPERCASE | `USERNAME` | Rare | Rare | ❌ | Sometimes |
| Hungarian Notation | `strName` | Legacy | Legacy | ❌ | ❌ |
| Prefix Notation | `_userName` | Framework Specific | Framework Specific | ❌ | ❌ |
| Suffix Notation | `userNameStr` | Rare | Rare | ❌ | ❌ |

---

# Go Naming Convention

Go follows a very simple naming style.

| Identifier | Convention | Example |
|------------|------------|---------|
| Package | lowercase | `package calculator` |
| Local Variable | camelCase | `userName` |
| Function | camelCase | `calculateTax()` |
| Exported Function | PascalCase | `CalculateTax()` |
| Struct | PascalCase | `User` |
| Interface | PascalCase | `Reader` |
| Constant | PascalCase or ALL_CAPS (rare) | `MaxUsers`, `PI` |

### Example

```go
package calculator

const MaxUsers = 100

type User struct {
    Name string
}

func CalculateTotal(price float64) float64 {
    return price
}

func calculateTax(price float64) float64 {
    return price * 0.075
}
```

---

# Best Practices

| Do ✅ | Don't ❌ |
|-------|----------|
| Use meaningful names | Use single-letter names unnecessarily |
| Follow your language's convention | Mix naming styles randomly |
| Keep names descriptive | Use abbreviations nobody understands |
| Be consistent | Switch conventions within the same project |
| Use camelCase or PascalCase where appropriate | Use Hungarian notation unless required |

---

# Quick Cheat Sheet

| Convention | Example |
|------------|---------|
| camelCase | `userName` |
| PascalCase | `UserName` |
| snake_case | `user_name` |
| SCREAMING_SNAKE_CASE | `USER_NAME` |
| kebab-case | `user-name` |
| Train-Case | `User-Name` |
| flatcase | `username` |
| UPPERCASE | `USERNAME` |
| Hungarian Notation | `strUserName` |
| Prefix Notation | `_userName` |
| Suffix Notation | `userNameStr` |

---

> **Rule of Thumb**
>
> - **Go:** camelCase + PascalCase
> - **Python:** snake_case
> - **JavaScript:** camelCase
> - **Java:** camelCase + PascalCase
> - **C#:** PascalCase + camelCase
> - **Rust:** snake_case
> - **CSS:** kebab-case
> - **URLs:** kebab-case
> - **Environment Variables:** SCREAMING_SNAKE_CASE