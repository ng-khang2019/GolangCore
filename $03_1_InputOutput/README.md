# Input and Output in Golang

## Introduction

Input and Output (I/O) are basic parts of programming.

- Input → receiving data from the user
- Output → displaying data to the screen

In Golang, most input and output operations use the `fmt` package.

---

# Importing the `fmt` Package

Example:

```go
import "fmt"
```

The `fmt` package provides functions for printing and reading data.

---

# Output in Golang

## Using `fmt.Print()`

Prints output without a new line.

Example:

```go
package main

import "fmt"

func main() {
    fmt.Print("Hello ")
    fmt.Print("World")
}
```

Output:

```text
Hello World
```

---

# Using `fmt.Println()`

Prints output with a new line.

Example:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello")
    fmt.Println("World")
}
```

Output:

```text
Hello
World
```

---

# Using `fmt.Printf()`

Used for formatted output.

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
| `%f` | Float |
| `%s` | String |
| `%t` | Boolean |
| `%v` | Default format |

Example:

```go
fmt.Printf("%d\n", 100)
fmt.Printf("%f\n", 3.14)
fmt.Printf("%s\n", "Hello")
```

---

# Input in Golang

## Using `fmt.Scan()`

Reads user input.

Example:

```go
package main

import "fmt"

func main() {
    var name string

    fmt.Print("Enter your name: ")
    fmt.Scan(&name)

    fmt.Println("Hello", name)
}
```

Important:

- `&name` passes the memory address
- Go stores the input value inside the variable

---

# Reading Multiple Inputs

Example:

```go
package main

import "fmt"

func main() {
    var name string
    var age int

    fmt.Scan(&name, &age)

    fmt.Println(name)
    fmt.Println(age)
}
```

Input:

```text
John 25
```

---

# Using `fmt.Scanln()`

Reads input until the Enter key is pressed.

Example:

```go
var city string

fmt.Scanln(&city)
```

---

# Using `fmt.Sprintf()`

Creates a formatted string instead of printing it.

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
    var name string
    var age int

    fmt.Print("Enter your name: ")
    fmt.Scan(&name)

    fmt.Print("Enter your age: ")
    fmt.Scan(&age)

    fmt.Printf("Name: %s\n", name)
    fmt.Printf("Age: %d\n", age)
}
```

---

# Common Beginner Mistakes

## Forgetting `&` in Input

Incorrect:

```go
fmt.Scan(name)
```

Correct:

```go
fmt.Scan(&name)
```

---

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

# Summary

Important points about input and output in Go:

- The `fmt` package is commonly used for I/O
- `Print()` prints without a new line
- `Println()` prints with a new line
- `Printf()` provides formatted output
- `Scan()` reads user input
- `&` is required when reading input into variables
- Format specifiers help display different data types

Input and output are essential for interacting with users in Golang applications. with a very specific format.