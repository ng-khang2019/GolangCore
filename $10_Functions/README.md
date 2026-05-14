# Functions in Golang

## Introduction

A function is a reusable block of code that performs a specific task.

Functions help:

- Organize code
- Reduce repetition
- Improve readability
- Make programs easier to maintain

Example:

```go
func greet() {
    fmt.Println("Hello")
}
```

---

# Basic Function Syntax

Syntax:

```go
func functionName() {
    // code
}
```

Example:

```go
package main

import "fmt"

func greet() {
    fmt.Println("Hello Golang")
}

func main() {
    greet()
}
```

Output:

```text
Hello Golang
```

---

# Function Parameters

Functions can receive input values called parameters.

Syntax:

```go
func functionName(parameter type) {
    // code
}
```

Example:

```go
package main

import "fmt"

func greet(name string) {
    fmt.Println("Hello", name)
}

func main() {
    greet("Alice")
}
```

Output:

```text
Hello Alice
```

---

# Multiple Parameters

Example:

```go
func add(a int, b int) {
    fmt.Println(a + b)
}
```

Calling the function:

```go
add(10, 5)
```

---

# Short Parameter Type Syntax

If multiple parameters share the same type:

```go
func add(a, b int) {
    fmt.Println(a + b)
}
```

This is commonly used in Go.

---

# Return Values

Functions can return values.

Syntax:

```go
func functionName() returnType {
    return value
}
```

Example:

```go
package main

import "fmt"

func add(a, b int) int {
    return a + b
}

func main() {
    result := add(10, 5)

    fmt.Println(result)
}
```

Output:

```text
15
```

---

# Returning Multiple Values

Go functions can return multiple values.

Example:

```go
package main

import "fmt"

func divide(a, b int) (int, int) {
    quotient := a / b
    remainder := a % b

    return quotient, remainder
}

func main() {
    q, r := divide(10, 3)

    fmt.Println(q)
    fmt.Println(r)
}
```

Output:

```text
3
1
```

---

# Ignoring Return Values

Use `_` to ignore unused return values.

Example:

```go
q, _ := divide(10, 3)
```

---

# Named Return Values

Go allows named return variables.

Example:

```go
func add(a, b int) (result int) {
    result = a + b
    return
}
```

This style is less common for simple functions.

---

# Function Scope

Variables declared inside a function are local to that function.

Example:

```go
func test() {
    x := 10
}
```

`x` cannot be accessed outside the function.

---

# Recursive Functions

A recursive function calls itself.

Example:

```go
func countdown(n int) {
    if n == 0 {
        return
    }

    fmt.Println(n)
    countdown(n - 1)
}
```

---

# Anonymous Functions

Functions without names are called anonymous functions.

Example:

```go
func() {
    fmt.Println("Anonymous Function")
}()
```

---

# Example Program

```go
package main

import "fmt"

func multiply(a, b int) int {
    return a * b
}

func main() {
    result := multiply(4, 5)

    fmt.Println("Result:", result)
}
```

Output:

```text
Result: 20
```

---

# Common Beginner Mistakes

## Forgetting Return Type

Incorrect:

```go
func add(a, b int) {
    return a + b
}
```

Correct:

```go
func add(a, b int) int {
    return a + b
}
```

---

## Wrong Number of Arguments

Incorrect:

```go
add(10)
```

Correct:

```go
add(10, 5)
```

---

# Summary

Important points about functions in Go:

- Functions are reusable blocks of code
- `func` is used to declare functions
- Functions can receive parameters
- Functions can return values
- Go supports multiple return values
- Variables inside functions are local
- Recursive and anonymous functions are supported
- Functions improve code organization and readability

Functions are one of the most important building blocks in Golang programming.