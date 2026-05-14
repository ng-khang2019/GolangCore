# Package in Golang

## What is a Package?

In Golang, a package is a way to organize and reuse code.  
Every Go program is made up of packages.

A package can contain:

- Functions
- Variables
- Constants
- Structs
- Interfaces
- Other Go files

Packages help keep code clean, modular, and easier to maintain.

---

# Basic Package Structure

Example:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello Golang")
}
```

Explanation:

- `package main` → defines the package name
- `import "fmt"` → imports another package
- `fmt.Println()` → uses a function from the `fmt` package

---

# The `main` Package

The `main` package is special.

A Go program starts running from:

```go
func main()
```

Example:

```go
package main

import "fmt"

func main() {
    fmt.Println("Program started")
}
```

Without the `main` package and `main()` function, the program cannot run as an executable application.

---

# Built-in Standard Packages

Go provides many built-in packages.

Some common packages:

| Package | Purpose |
|---|---|
| `fmt` | Input and output |
| `math` | Mathematical functions |
| `strings` | String processing |
| `time` | Date and time |
| `os` | Operating system features |
| `strconv` | String conversion |

Example:

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(math.Sqrt(25))
}
```

---

# Creating Your Own Package

Example folder structure:

```text
project/
│
├── main.go
└── calculator/
    └── calculator.go
```

---

## calculator/calculator.go

```go
package calculator

func Add(a int, b int) int {
    return a + b
}
```

---

## main.go

```go
package main

import (
    "fmt"
    "project/calculator"
)

func main() {
    result := calculator.Add(10, 5)
    fmt.Println(result)
}
```

---

# Exported vs Unexported Names

In Go:

- Names starting with an uppercase letter are exported
- Names starting with a lowercase letter are private to the package

Example:

```go
package example

func PublicFunction() {
}

func privateFunction() {
}
```

- `PublicFunction()` → accessible from other packages
- `privateFunction()` → only accessible inside the same package

---

# Importing Multiple Packages

Example:

```go
import (
    "fmt"
    "math"
    "strings"
)
```

This style is commonly used in Go projects.

---

# Package Naming Rules

Good package names should:

- Be short
- Be lowercase
- Be meaningful
- Avoid underscores and spaces

Good examples:

```text
math
utils
database
service
config
```

Bad examples:

```text
MyPackage
my_package
PACKAGE
```

---

# Why Packages Are Important

Packages help:

- Organize code
- Reuse code
- Improve readability
- Reduce duplication
- Make large projects manageable

Without packages, large applications become difficult to maintain.

---

# Common Go Project Structure

Example:

```text
project/
│
├── main.go
├── config/
├── models/
├── services/
├── handlers/
└── utils/
```

Each folder is usually a separate package.

---

# Summary

A package in Go is a collection of related code.

Important points:

- Every Go file belongs to a package
- `main` is the executable package
- Packages help organize code
- Use `import` to use other packages
- Uppercase names are exported
- Lowercase names are private

Packages are one of the core concepts in Golang and are essential for building clean and scalable applications.