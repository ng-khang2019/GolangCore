# Common String Functions in Golang

## Introduction

Golang provides many useful string functions through the `strings` package.

Before using these functions, import the package:

```go
import "strings"
```

These functions help with:

- Searching text
- Replacing text
- Splitting strings
- Formatting text
- Comparing strings

---

# `strings.Contains()`

Checks whether a string contains another string.

Syntax:

```go
strings.Contains(string, substring)
```

Example:

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    result := strings.Contains("Hello Golang", "Go")

    fmt.Println(result)
}
```

Output:

```text
true
```

---

# `strings.HasPrefix()`

Checks whether a string starts with a specific prefix.

Example:

```go
fmt.Println(strings.HasPrefix("Golang", "Go"))
```

Output:

```text
true
```

---

# `strings.HasSuffix()`

Checks whether a string ends with a specific suffix.

Example:

```go
fmt.Println(strings.HasSuffix("file.txt", ".txt"))
```

Output:

```text
true
```

---

# `strings.ToUpper()`

Converts all characters to uppercase.

Example:

```go
fmt.Println(strings.ToUpper("golang"))
```

Output:

```text
GOLANG
```

---

# `strings.ToLower()`

Converts all characters to lowercase.

Example:

```go
fmt.Println(strings.ToLower("HELLO"))
```

Output:

```text
hello
```

---

# `strings.TrimSpace()`

Removes spaces at the beginning and end of a string.

Example:

```go
text := "   Hello Go   "

fmt.Println(strings.TrimSpace(text))
```

Output:

```text
Hello Go
```

---

# `strings.ReplaceAll()`

Replaces all matching substrings.

Syntax:

```go
strings.ReplaceAll(string, old, new)
```

Example:

```go
fmt.Println(strings.ReplaceAll("Go Go Go", "Go", "Hi"))
```

Output:

```text
Hi Hi Hi
```

---

# `strings.Split()`

Splits a string into a slice.

Syntax:

```go
strings.Split(string, separator)
```

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

# `strings.Join()`

Joins slice elements into a single string.

Syntax:

```go
strings.Join(slice, separator)
```

Example:

```go
fruits := []string{"apple", "banana", "orange"}

result := strings.Join(fruits, "-")

fmt.Println(result)
```

Output:

```text
apple-banana-orange
```

---

# `strings.Index()`

Finds the position of a substring.

Example:

```go
fmt.Println(strings.Index("Golang", "lang"))
```

Output:

```text
2
```

If not found:

```text
-1
```

---

# `strings.Count()`

Counts occurrences of a substring.

Example:

```go
fmt.Println(strings.Count("Go Go Go", "Go"))
```

Output:

```text
3
```

---

# `strings.Repeat()`

Repeats a string multiple times.

Example:

```go
fmt.Println(strings.Repeat("Go ", 3))
```

Output:

```text
Go Go Go
```

---

# `strings.Compare()`

Compares two strings.

Example:

```go
fmt.Println(strings.Compare("abc", "abc"))
fmt.Println(strings.Compare("abc", "xyz"))
```

Possible outputs:

| Result | Meaning |
|---|---|
| `0` | Equal |
| `-1` | First string is smaller |
| `1` | First string is greater |

---

# `strings.Fields()`

Splits a string by spaces.

Example:

```go
text := "Go is fun"

result := strings.Fields(text)

fmt.Println(result)
```

Output:

```text
[Go is fun]
```

---

# `strings.ContainsAny()`

Checks whether a string contains any character from another string.

Example:

```go
fmt.Println(strings.ContainsAny("Golang", "xyzG"))
```

Output:

```text
true
```

---

# `strings.EqualFold()`

Compares strings without case sensitivity.

Example:

```go
fmt.Println(strings.EqualFold("Go", "go"))
```

Output:

```text
true
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
    text := "  golang programming  "

    fmt.Println(strings.TrimSpace(text))
    fmt.Println(strings.ToUpper(text))
    fmt.Println(strings.Contains(text, "go"))
    fmt.Println(strings.ReplaceAll(text, "golang", "Go"))
}
```

---

# Common Beginner Mistakes

## Forgetting to Import `strings`

Incorrect:

```go
strings.ToUpper("hello")
```

Without:

```go
import "strings"
```

the program will fail.

---

## Case Sensitivity

Incorrect expectation:

```go
strings.Contains("GoLang", "go")
```

Output:

```text
false
```

String functions are usually case-sensitive.

---

# Summary

Important points about string functions in Go:

- The `strings` package provides many useful utilities
- `Contains()` searches text
- `ToUpper()` and `ToLower()` change case
- `Split()` separates strings
- `Join()` combines strings
- `ReplaceAll()` replaces text
- `TrimSpace()` removes extra spaces
- `Index()` finds substring positions
- Most functions are case-sensitive

String functions are essential for text processing in Golang applications.