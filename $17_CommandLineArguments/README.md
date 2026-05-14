# Command Line Arguments in Golang

## Introduction

Command Line Arguments are values passed to a program when it starts running.

They allow users to provide input directly from the terminal.

Example:

```bash
go run main.go Alice 25
```

Here:

- `Alice`
- `25`

are command line arguments.

---

# Why Use Command Line Arguments?

Command line arguments help:

- Pass user input quickly
- Configure programs
- Automate tasks
- Build command-line tools

They are commonly used in real-world applications.

---

# The `os` Package

Go uses the `os` package to access command line arguments.

Import:

```go
import "os"
```

Arguments are stored in:

```go
os.Args
```

---

# Understanding `os.Args`

`os.Args` is a slice of strings.

Example:

```go
fmt.Println(os.Args)
```

If the command is:

```bash
go run main.go hello world
```

Output may look like:

```text
[/tmp/main hello world]
```

Important:

- `os.Args[0]` → program name
- Actual arguments start from `os.Args[1]`

---

# Accessing Arguments

Example:

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("First Argument:", os.Args[1])
}
```

Run:

```bash
go run main.go Golang
```

Output:

```text
First Argument: Golang
```

---

# Reading Multiple Arguments

Example:

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    fmt.Println("Name:", os.Args[1])
    fmt.Println("Age:", os.Args[2])
}
```

Run:

```bash
go run main.go Alice 22
```

Output:

```text
Name: Alice
Age: 22
```

---

# Looping Through Arguments

Use `range` to loop through all arguments.

Example:

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    for index, value := range os.Args {
        fmt.Println(index, value)
    }
}
```

---

# Checking Argument Length

Always check the number of arguments before accessing them.

Example:

```go
if len(os.Args) < 2 {
    fmt.Println("Please provide an argument")
    return
}
```

This prevents runtime errors.

---

# Converting Arguments to Numbers

Command line arguments are strings by default.

Use `strconv` to convert them.

Example:

```go
package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    number, err := strconv.Atoi(os.Args[1])

    if err != nil {
        fmt.Println("Invalid number")
        return
    }

    fmt.Println(number * 2)
}
```

Run:

```bash
go run main.go 10
```

Output:

```text
20
```

---

# Example Program

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: go run main.go <name> <city>")
        return
    }

    name := os.Args[1]
    city := os.Args[2]

    fmt.Println("Name:", name)
    fmt.Println("City:", city)
}
```

Run:

```bash
go run main.go Alice Tokyo
```

Output:

```text
Name: Alice
City: Tokyo
```

---

# Common Beginner Mistakes

## Accessing Missing Arguments

Incorrect:

```go
fmt.Println(os.Args[1])
```

without checking length.

This may cause:

```text
panic: runtime error
```

Correct:

```go
if len(os.Args) > 1 {
    fmt.Println(os.Args[1])
}
```

---

## Forgetting Arguments Are Strings

Incorrect:

```go
result := os.Args[1] + 10
```

Correct:

```go
number, _ := strconv.Atoi(os.Args[1])
```

---

# Summary

Important points about command line arguments in Go:

- Arguments are passed through the terminal
- `os.Args` stores command line arguments
- `os.Args[0]` is the program name
- Arguments are stored as strings
- `len()` helps prevent errors
- `range` can loop through arguments
- `strconv` converts strings to numbers

Command line arguments are very useful for building flexible and interactive Golang programs.