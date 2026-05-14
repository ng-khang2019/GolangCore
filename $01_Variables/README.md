# Variables in Golang

## What is a Variable?

A variable is a container used to store data in a program.

In Golang, variables can store different types of values such as:

- Numbers
- Text
- Boolean values
- Arrays
- Structs

Example:

```go
var name string = "John"
```

Here:

- `name` → variable name
- `string` → data type
- `"John"` → stored value

---

# Declaring Variables

## Using `var`

Syntax:

```go
var variableName dataType = value
```

Example:

```go
var age int = 25
var price float64 = 99.99
var isOnline bool = true
```

---

# Type Inference

Go can automatically detect the data type.

Example:

```go
var city = "Tokyo"
var number = 100
```

Go automatically understands:

- `"Tokyo"` → string
- `100` → int

---

# Short Variable Declaration

Inside functions, you can use `:=`

Example:

```go
name := "Alice"
score := 95
```

This is the most commonly used style in Go.

Important:

- `:=` can only be used inside functions
- It cannot be used outside functions

---

# Multiple Variable Declaration

Example:

```go
var a, b, c int = 1, 2, 3
```

Or:

```go
x, y := 10, 20
```

---

# Variable Data Types

Common data types in Go:

| Type | Description |
|---|---|
| `int` | Integer numbers |
| `float64` | Decimal numbers |
| `string` | Text |
| `bool` | true or false |

Example:

```go
var age int = 30
var height float64 = 175.5
var name string = "Mike"
var isStudent bool = false
```

---

# Zero Values

If a variable is declared without a value, Go gives it a default value.

Examples:

| Type | Zero Value |
|---|---|
| `int` | `0` |
| `float64` | `0` |
| `string` | `""` |
| `bool` | `false` |

Example:

```go
var number int
fmt.Println(number)
```

Output:

```text
0
```

---

# Constants

Constants are values that cannot be changed.

Syntax:

```go
const PI = 3.14
```

Example:

```go
const country = "Japan"
```

Trying to change a constant will cause an error.

---

# Variable Scope

## Local Variables

Declared inside a function.

Example:

```go
func main() {
    name := "John"
}
```

Only accessible inside that function.

---

## Global Variables

Declared outside functions.

Example:

```go
package main

import "fmt"

var message = "Hello"

func main() {
    fmt.Println(message)
}
```

Accessible from multiple functions in the same package.

---

# Naming Rules

Variable names:

- Can contain letters, numbers, and `_`
- Cannot start with a number
- Are case-sensitive

Good examples:

```go
userName
totalPrice
studentAge
```

Bad examples:

```go
123name
total-price
```

---

# Common Naming Style in Go

Go commonly uses camelCase.

Example:

```go
firstName
totalAmount
isLoggedIn
```

---

# Example Program

```go
package main

import "fmt"

func main() {
    name := "Alice"
    age := 22
    height := 165.5
    isStudent := true

    fmt.Println(name)
    fmt.Println(age)
    fmt.Println(height)
    fmt.Println(isStudent)
}
```

---

# Summary

Important points about variables in Go:

- Variables store data
- `var` is the standard declaration keyword
- `:=` is shorthand declaration inside functions
- Go supports type inference
- Variables have data types
- Uninitialized variables receive zero values
- Constants cannot be changed
- Variable names should follow Go naming conventions

Variables are one of the most fundamental concepts in Golang programming.