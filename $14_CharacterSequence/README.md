# Strings and Character Sequences in Golang

## Introduction

A string in Golang is a sequence of characters used to store text.

Examples:

```go
"Hello"
"Go Programming"
"123"
```

Strings are widely used for:

- User input
- Messages
- File names
- Text processing
- Data storage

---

# Declaring Strings

## Using `var`

Example:

```go
var name string = "Alice"
```

---

## Using Short Declaration

Example:

```go
message := "Hello Golang"
```

---

# String Characteristics

Important points about strings in Go:

- Strings are immutable
- Strings are sequences of bytes
- String indexes start at `0`
- Strings support Unicode

Immutable means the original string cannot be directly changed.

---

# Accessing Characters

You can access characters using indexes.

Example:

```go
text := "Hello"

fmt.Println(text[0])
fmt.Println(text[1])
```

Output:

```text
72
101
```

Important:

Go returns ASCII/byte values when accessing string indexes.

---

# Converting Byte to Character

Example:

```go
text := "Hello"

fmt.Printf("%c\n", text[0])
```

Output:

```text
H
```

---

# String Length

Use `len()` to get the string length.

Example:

```go
text := "Hello"

fmt.Println(len(text))
```

Output:

```text
5
```

---

# String Concatenation

Strings can be combined using `+`.

Example:

```go
firstName := "John"
lastName := "Doe"

fullName := firstName + " " + lastName

fmt.Println(fullName)
```

Output:

```text
John Doe
```

---

# Multi-line Strings

Use backticks `` ` ` `` for multi-line strings.

Example:

```go
text := `Line 1
Line 2
Line 3`

fmt.Println(text)
```

---

# Escape Characters

| Escape Character | Description |
|---|---|
| `\n` | New line |
| `\t` | Tab |
| `\"` | Double quote |
| `\\` | Backslash |

Example:

```go
fmt.Println("Hello\nWorld")
```

Output:

```text
Hello
World
```

---

# Looping Through Strings

## Using Traditional `for`

Example:

```go
text := "Go"

for i := 0; i < len(text); i++ {
    fmt.Printf("%c\n", text[i])
}
```

---

# Using `range`

`range` is recommended for Unicode strings.

Example:

```go
text := "Hello"

for index, char := range text {
    fmt.Println(index, string(char))
}
```

---

# Unicode and Runes

Go uses `rune` for Unicode characters.

Example:

```go
char := 'A'

fmt.Println(char)
```

Output:

```text
65
```

---

# Converting String to Rune Slice

Example:

```go
text := "Hello"

runes := []rune(text)

fmt.Println(runes)
```

This is useful when working with Unicode characters.

---

# Common String Functions

The `strings` package provides many useful functions.

Import:

```go
import "strings"
```

---

## `strings.Contains()`

Checks if a string contains another string.

Example:

```go
fmt.Println(strings.Contains("Hello", "He"))
```

---

## `strings.ToUpper()`

Converts to uppercase.

Example:

```go
fmt.Println(strings.ToUpper("hello"))
```

Output:

```text
HELLO
```

---

## `strings.ToLower()`

Converts to lowercase.

Example:

```go
fmt.Println(strings.ToLower("HELLO"))
```

---

## `strings.ReplaceAll()`

Replaces text.

Example:

```go
fmt.Println(strings.ReplaceAll("Go Go", "Go", "Hi"))
```

Output:

```text
Hi Hi
```

---

## `strings.Split()`

Splits a string into slices.

Example:

```go
data := "apple,banana,orange"

result := strings.Split(data, ",")

fmt.Println(result)
```

Output:

```text
[apple banana orange]
```

---

# Strings Are Immutable

Incorrect:

```go
text := "Hello"

text[0] = 'A'
```

This causes an error.

Correct approach:

Convert to a rune slice.

Example:

```go
text := "Hello"

runes := []rune(text)

runes[0] = 'A'

text = string(runes)

fmt.Println(text)
```

Output:

```text
Aello
```

---

# Example Program

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    text := "golang"

    fmt.Println(strings.ToUpper(text))
    fmt.Println(len(text))

    for _, char := range text {
        fmt.Printf("%c ", char)
    }
}
```

Output:

```text
GOLANG
6
g o l a n g
```

---

# Common Beginner Mistakes

## Confusing Characters and Strings

Incorrect:

```go
char := "A"
```

This is a string, not a rune.

Correct:

```go
char := 'A'
```

---

## Trying to Modify Strings Directly

Incorrect:

```go
text[0] = 'H'
```

Strings are immutable.

---

# Summary

Important points about strings in Go:

- Strings store text data
- Strings are immutable
- Indexes start at `0`
- `len()` returns string length
- `+` concatenates strings
- `range` is useful for iterating through strings
- Go supports Unicode using `rune`
- The `strings` package provides many useful functions
- Escape characters help format text

Strings are one of the most important data types in Golang programming.