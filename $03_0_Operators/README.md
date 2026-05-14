# Operators in Golang

## What is an Operator?

An operator is a symbol used to perform operations on variables and values.

Example:

```go
a + b
```

Here:

- `+` is an operator
- `a` and `b` are operands

---

# Types of Operators in Go

Go provides several types of operators:

- Arithmetic Operators
- Assignment Operators
- Comparison Operators
- Logical Operators
- Increment and Decrement Operators
- Bitwise Operators

---

# Arithmetic Operators

Used for mathematical calculations.

| Operator | Description | Example |
|---|---|---|
| `+` | Addition | `a + b` |
| `-` | Subtraction | `a - b` |
| `*` | Multiplication | `a * b` |
| `/` | Division | `a / b` |
| `%` | Modulus | `a % b` |

Example:

```go
package main

import "fmt"

func main() {
    a := 10
    b := 3

    fmt.Println(a + b)
    fmt.Println(a - b)
    fmt.Println(a * b)
    fmt.Println(a / b)
    fmt.Println(a % b)
}
```

---

# Assignment Operators

Used to assign values to variables.

| Operator | Example | Meaning |
|---|---|---|
| `=` | `x = 5` | Assign value |
| `+=` | `x += 2` | `x = x + 2` |
| `-=` | `x -= 2` | `x = x - 2` |
| `*=` | `x *= 2` | `x = x * 2` |
| `/=` | `x /= 2` | `x = x / 2` |

Example:

```go
x := 10

x += 5
fmt.Println(x)
```

Output:

```text
15
```

---

# Comparison Operators

Used to compare values.

| Operator | Description |
|---|---|
| `==` | Equal to |
| `!=` | Not equal to |
| `>` | Greater than |
| `<` | Less than |
| `>=` | Greater than or equal to |
| `<=` | Less than or equal to |

Example:

```go
a := 10
b := 20

fmt.Println(a == b)
fmt.Println(a < b)
fmt.Println(a != b)
```

Comparison operators return boolean values:

```text
true
false
```

---

# Logical Operators

Used with boolean values.

| Operator | Description |
|---|---|
| `&&` | Logical AND |
| `||` | Logical OR |
| `!` | Logical NOT |

Example:

```go
age := 20
hasTicket := true

fmt.Println(age >= 18 && hasTicket)
```

---

# Increment and Decrement Operators

Used to increase or decrease values.

| Operator | Description |
|---|---|
| `++` | Increment |
| `--` | Decrement |

Example:

```go
count := 5

count++
fmt.Println(count)

count--
fmt.Println(count)
```

Important:

In Go, `++` and `--` are statements, not expressions.

Correct:

```go
count++
```

Incorrect:

```go
x = count++
```

---

# Bitwise Operators

Used for binary operations.

| Operator | Description |
|---|---|
| `&` | AND |
| `|` | OR |
| `^` | XOR |
| `<<` | Left shift |
| `>>` | Right shift |

Example:

```go
a := 5
b := 3

fmt.Println(a & b)
fmt.Println(a | b)
```

Bitwise operators are commonly used in low-level programming.

---

# Operator Precedence

Some operators execute before others.

Example:

```go
result := 10 + 5 * 2
```

Multiplication happens first:

```text
10 + (5 * 2) = 20
```

Use parentheses for clarity:

```go
result := (10 + 5) * 2
```

Output:

```text
30
```

---

# Example Program

```go
package main

import "fmt"

func main() {
    a := 10
    b := 5

    fmt.Println(a + b)
    fmt.Println(a > b)
    fmt.Println(a == b)

    isValid := true
    isAdmin := false

    fmt.Println(isValid && isAdmin)
}
```

---

# Summary

Important points about operators in Go:

- Operators perform actions on values and variables
- Arithmetic operators are used for math
- Assignment operators update variable values
- Comparison operators return boolean results
- Logical operators work with boolean values
- `++` and `--` are statements in Go
- Bitwise operators work with binary data
- Operator precedence affects calculation order

Operators are essential for writing logic and calculations in Golang programs.