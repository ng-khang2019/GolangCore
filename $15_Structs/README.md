# Struct in Golang

## Introduction

A struct in Golang is a custom data type used to group related data together.

Structs are useful when storing multiple pieces of information about one object.

Example:

```go
type Person struct {
    Name string
    Age  int
}
```

This struct stores information about a person.

---

# Why Use Structs?

Structs help:

- Organize related data
- Create custom data types
- Represent real-world objects
- Improve code readability

Without structs, managing related data becomes difficult.

---

# Declaring a Struct

Syntax:

```go
type StructName struct {
    fieldName dataType
}
```

Example:

```go
type Student struct {
    Name  string
    Age   int
    Score float64
}
```

---

# Creating Struct Variables

## Method 1: Full Initialization

Example:

```go
student := Student{
    Name:  "Alice",
    Age:   20,
    Score: 95.5,
}
```

---

## Method 2: Positional Initialization

Example:

```go
student := Student{"Bob", 22, 88.5}
```

Important:

Values must follow the field order.

---

## Method 3: Empty Struct

Example:

```go
var student Student
```

Fields receive zero values.

---

# Accessing Struct Fields

Use the dot `.` operator.

Example:

```go
fmt.Println(student.Name)
fmt.Println(student.Age)
```

---

# Modifying Struct Fields

Example:

```go
student.Name = "John"
student.Age = 25
```

---

# Example Program

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    person := Person{
        Name: "Alice",
        Age:  22,
    }

    fmt.Println(person.Name)
    fmt.Println(person.Age)
}
```

Output:

```text
Alice
22
```

---

# Structs with Functions

Structs are often used with functions.

Example:

```go
type Rectangle struct {
    Width  float64
    Height float64
}

func area(r Rectangle) float64 {
    return r.Width * r.Height
}
```

---

# Anonymous Structs

Go supports structs without names.

Example:

```go
person := struct {
    Name string
    Age  int
}{
    Name: "Tom",
    Age:  30,
}
```

---

# Nested Structs

Structs can contain other structs.

Example:

```go
type Address struct {
    City string
}

type Person struct {
    Name    string
    Address Address
}
```

Accessing nested fields:

```go
fmt.Println(person.Address.City)
```

---

# Struct Comparison

Structs can be compared if all fields are comparable.

Example:

```go
a := Person{"John", 20}
b := Person{"John", 20}

fmt.Println(a == b)
```

Output:

```text
true
```

---

# Struct Pointers

Structs can be used with pointers.

Example:

```go
person := Person{"Alice", 22}

ptr := &person

fmt.Println(ptr.Name)
```

Go automatically dereferences struct pointers.

---

# Exported and Unexported Fields

Field names starting with uppercase letters are exported.

Example:

```go
type User struct {
    Name string
    age  int
}
```

- `Name` → accessible outside the package
- `age` → private inside the package

---

# Zero Values in Structs

Uninitialized struct fields receive zero values.

Example:

```go
var person Person

fmt.Println(person.Name)
fmt.Println(person.Age)
```

Output:

```text

0
```

---

# Struct Tags

Struct fields can contain tags.

Example:

```go
type User struct {
    Name string `json:"name"`
}
```

Tags are commonly used in JSON and database operations.

---

# Example with Slice of Structs

Example:

```go
students := []Student{
    {"Alice", 20, 90},
    {"Bob", 21, 85},
}
```

Looping:

```go
for _, student := range students {
    fmt.Println(student.Name)
}
```

---

# Common Beginner Mistakes

## Forgetting Field Names

Incorrect:

```go
student.Name()
```

Correct:

```go
student.Name
```

Fields are accessed without parentheses.

---

## Wrong Field Order

Incorrect:

```go
Student{20, "Alice", 90}
```

Field types/order must match.

---

# Summary

Important points about structs in Go:

- Structs group related data together
- `type` and `struct` are used to define structs
- Fields are accessed using `.`
- Structs support nested structures
- Structs can work with functions and pointers
- Uppercase fields are exported
- Structs receive zero values by default
- Struct tags are useful for JSON and databases

Structs are essential for building organized and scalable Golang applications.