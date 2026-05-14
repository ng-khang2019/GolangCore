# Formatting Output in Golang

## Introduction

Formatting output means displaying data in a specific and readable format.

In Golang, formatted output is mainly handled using the `fmt` package.

The most commonly used function is:

```go
fmt.Printf()
```

---

# Why Formatting Output is Important

Formatted output helps:

- Display data clearly
- Control how values appear
- Print variables with text
- Improve program readability

Example:

```go
fmt.Printf("Age: %d", 25)
```

Output:

```text
Age: 25
```

---

# Using `fmt.Printf()`

Syntax:

```go
fmt.Printf("format", values)
```

Example:

```go
package main

import "fmt"

func main() {
    name := "Alice"
    age := 22

    fmt.Printf("Name: %s\n", name)
    fmt.Printf("Age: %d\n", age)
}
```

Output:

```text
Name: Alice
Age: 22
```

---

# Common Format Specifiers

| Specifier | Description |
|---|---|
| `%d` | Integer |
| `%f` | Floating-point number |
| `%s` | String |
| `%t` | Boolean |
| `%c` | Character |
| `%v` | Default value format |
| `%%` | Percent sign |

---

# Integer Formatting

Example:

```go
number := 100

fmt.Printf("%d\n", number)
```

Output:

```text
100
```

---

# Float Formatting

Example:

```go
price := 19.99

fmt.Printf("%f\n", price)
```

Output:

```text
19.990000
```

---

# Limiting Decimal Places

Example:

```go
fmt.Printf("%.2f\n", price)
```

Output:

```text
19.99
```

Explanation:

- `.2` means display 2 digits after the decimal point

---

# String Formatting

Example:

```go
name := "John"

fmt.Printf("%s\n", name)
```

Output:

```text
John
```

---

# Boolean Formatting

Example:

```go
isOnline := true

fmt.Printf("%t\n", isOnline)
```

Output:

```text
true
```

---

# Printing Multiple Values

Example:

```go
name := "Alice"
age := 25

fmt.Printf("Name: %s, Age: %d\n", name, age)
```

Output:

```text
Name: Alice, Age: 25
```

---

# Using `%v`

`%v` displays values in their default format.

Example:

```go
fmt.Printf("%v\n", 100)
fmt.Printf("%v\n", "Hello")
fmt.Printf("%v\n", true)
```

---

# Printing a Percent Sign

Use `%%` to print `%`.

Example:

```go
fmt.Printf("Success Rate: 90%%\n")
```

Output:

```text
Success Rate: 90%
```

---

# Width and Alignment

You can control spacing and alignment.

Example:

```go
fmt.Printf("%10d\n", 50)
```

Output:

```text
        50
```

Explanation:

- `10` means minimum width of 10 characters

---

# Left Alignment

Example:

```go
fmt.Printf("%-10sEND\n", "Go")
```

Output:

```text
Go        END
```

---

# Using `fmt.Sprintf()`

`Sprintf()` formats output and returns it as a string.

Example:

```go
name := "Bob"

message := fmt.Sprintf("Hello %s", name)

fmt.Println(message)
```

---

# Example Program

```go
package main

import "fmt"

func main() {
    name := "Alice"
    age := 22
    score := 95.456

    fmt.Printf("Name : %s\n", name)
    fmt.Printf("Age  : %d\n", age)
    fmt.Printf("Score: %.2f\n", score)
}
```

Output:

```text
Name : Alice
Age  : 22
Score: 95.46
```

---

# Common Beginner Mistakes

## Using Wrong Format Specifiers

Incorrect:

```go
fmt.Printf("%d", "Hello")
```

Correct:

```go
fmt.Printf("%s", "Hello")
```

---

## Forgetting `\n`

Without `\n`, output may appear on the same line.

Example:

```go
fmt.Printf("Hello")
fmt.Printf("World")
```

Output:

```text
HelloWorld
```

---

# Summary

Important points about formatting output in Go:

- `fmt.Printf()` is used for formatted output
- Format specifiers control how values are displayed
- `%d`, `%f`, `%s`, and `%t` are commonly used
- `%.2f` limits decimal places
- `%v` displays values in default format
- `fmt.Sprintf()` returns formatted strings
- Width and alignment improve readability

Formatted output is very useful for displaying clean and professional program results.