# If-Else in Golang

## Introduction

`if-else` statements are used to make decisions in a program.

They allow the program to execute different blocks of code based on conditions.

Example:

```go
if age >= 18 {
    fmt.Println("Adult")
}
```

If the condition is true, the code inside the block runs.

---

# Basic `if` Statement

Syntax:

```go
if condition {
    // code
}
```

Example:

```go
package main

import "fmt"

func main() {
    age := 20

    if age >= 18 {
        fmt.Println("You are an adult")
    }
}
```

---

# `if-else` Statement

Used when there are two possible outcomes.

Syntax:

```go
if condition {
    // code if true
} else {
    // code if false
}
```

Example:

```go
package main

import "fmt"

func main() {
    age := 16

    if age >= 18 {
        fmt.Println("Adult")
    } else {
        fmt.Println("Minor")
    }
}
```

Output:

```text
Minor
```

---

# `else if` Statement

Used to check multiple conditions.

Syntax:

```go
if condition1 {
    // code
} else if condition2 {
    // code
} else {
    // code
}
```

Example:

```go
package main

import "fmt"

func main() {
    score := 85

    if score >= 90 {
        fmt.Println("Grade A")
    } else if score >= 80 {
        fmt.Println("Grade B")
    } else if score >= 70 {
        fmt.Println("Grade C")
    } else {
        fmt.Println("Grade D")
    }
}
```

Output:

```text
Grade B
```

---

# Comparison Operators in Conditions

Common operators used in conditions:

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
if number != 0 {
    fmt.Println("Not zero")
}
```

---

# Logical Operators in Conditions

Logical operators combine multiple conditions.

| Operator | Description |
|---|---|
| `&&` | AND |
| `||` | OR |
| `!` | NOT |

Example:

```go
age := 25
hasTicket := true

if age >= 18 && hasTicket {
    fmt.Println("Entry allowed")
}
```

---

# Nested `if` Statements

An `if` statement inside another `if` statement is called a nested `if`.

Example:

```go
age := 20
hasID := true

if age >= 18 {
    if hasID {
        fmt.Println("Access granted")
    }
}
```

---

# Short Statement in `if`

Go allows variable declaration inside `if`.

Syntax:

```go
if variable := value; condition {
    // code
}
```

Example:

```go
if score := 85; score >= 80 {
    fmt.Println("Passed")
}
```

The variable only exists inside the `if` block.

---

# Important Syntax Rules

## No Parentheses Needed

Correct:

```go
if age > 18 {
}
```

Incorrect:

```go
if (age > 18) {
}
```

---

## Curly Braces Are Required

Correct:

```go
if true {
    fmt.Println("Hello")
}
```

Incorrect:

```go
if true
    fmt.Println("Hello")
```

---

# Example Program

```go
package main

import "fmt"

func main() {
    temperature := 30

    if temperature > 35 {
        fmt.Println("Very Hot")
    } else if temperature >= 25 {
        fmt.Println("Warm")
    } else {
        fmt.Println("Cold")
    }
}
```

Output:

```text
Warm
```

---

# Common Beginner Mistakes

## Using `=` Instead of `==`

Incorrect:

```go
if age = 18 {
}
```

Correct:

```go
if age == 18 {
}
```

---

## Forgetting Curly Braces

Incorrect:

```go
if age > 18
    fmt.Println("Adult")
```

Go requires `{}`.

---

# Summary

Important points about `if-else` in Go:

- `if` executes code when a condition is true
- `else` handles the false case
- `else if` checks multiple conditions
- Conditions use comparison and logical operators
- Parentheses are not required in Go conditions
- Curly braces are mandatory
- Nested `if` statements are supported
- Variables can be declared inside `if` statements

`if-else` statements are essential for controlling program flow and decision-making in Golang.