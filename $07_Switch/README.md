# Switch in Golang

## Introduction

A `switch` statement is used to select one block of code from multiple possible options.

It is often cleaner and easier to read than many `if-else if` statements.

Example:

```go
switch day {
case 1:
    fmt.Println("Monday")
case 2:
    fmt.Println("Tuesday")
}
```

---

# Basic `switch` Syntax

Syntax:

```go
switch expression {
case value1:
    // code
case value2:
    // code
default:
    // code
}
```

Example:

```go
package main

import "fmt"

func main() {
    day := 2

    switch day {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    case 3:
        fmt.Println("Wednesday")
    default:
        fmt.Println("Unknown day")
    }
}
```

Output:

```text
Tuesday
```

---

# How `switch` Works

The `switch` expression is compared against each `case`.

- If a match is found, that block runs
- If no match exists, the `default` block runs

---

# The `default` Case

`default` works like `else`.

Example:

```go
number := 10

switch number {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
default:
    fmt.Println("Other number")
}
```

Output:

```text
Other number
```

---

# Multiple Values in One Case

A case can contain multiple values.

Example:

```go
letter := "a"

switch letter {
case "a", "e", "i", "o", "u":
    fmt.Println("Vowel")
default:
    fmt.Println("Consonant")
}
```

---

# Switch Without an Expression

Go allows `switch` without an expression.

This works similarly to multiple `if-else` conditions.

Example:

```go
score := 85

switch {
case score >= 90:
    fmt.Println("Grade A")
case score >= 80:
    fmt.Println("Grade B")
case score >= 70:
    fmt.Println("Grade C")
default:
    fmt.Println("Grade D")
}
```

Output:

```text
Grade B
```

---

# Important Difference from Other Languages

In Go, `break` is automatic.

Example:

```go
switch number {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
}
```

Go automatically exits the `switch` after a matched case.

You usually do not need to write:

```go
break
```

---

# Using `fallthrough`

`fallthrough` forces execution into the next case.

Example:

```go
number := 1

switch number {
case 1:
    fmt.Println("One")
    fallthrough
case 2:
    fmt.Println("Two")
}
```

Output:

```text
One
Two
```

Important:

- `fallthrough` ignores the next case condition
- It simply continues to the next block

---

# Short Statement in `switch`

Go allows variable declaration inside `switch`.

Example:

```go
switch day := 3; day {
case 1:
    fmt.Println("Monday")
case 2:
    fmt.Println("Tuesday")
case 3:
    fmt.Println("Wednesday")
}
```

The variable only exists inside the `switch`.

---

# Example Program

```go
package main

import "fmt"

func main() {
    month := 4

    switch month {
    case 12, 1, 2:
        fmt.Println("Winter")
    case 3, 4, 5:
        fmt.Println("Spring")
    case 6, 7, 8:
        fmt.Println("Summer")
    case 9, 10, 11:
        fmt.Println("Autumn")
    default:
        fmt.Println("Invalid month")
    }
}
```

Output:

```text
Spring
```

---

# Common Beginner Mistakes

## Forgetting `default`

Not required, but useful for handling unexpected values.

---

## Misusing `fallthrough`

Incorrect use of `fallthrough` can cause unexpected behavior.

Example:

```go
fallthrough
```

always runs the next case block.

---

# Summary

Important points about `switch` in Go:

- `switch` selects code based on matching cases
- `default` runs when no case matches
- Multiple values can exist in one case
- `switch` can be used without an expression
- Go automatically breaks after each case
- `fallthrough` forces execution into the next case
- Variables can be declared inside `switch`

`switch` statements are useful for writing clean and readable decision-making logic in Golang.