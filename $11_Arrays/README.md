# Arrays in Golang

## Introduction

An array is a collection of elements of the same data type stored in a fixed size.

Example:

```go
var numbers [5]int
```

This creates an array that can store 5 integers.

Arrays are useful when you need to store multiple values together.

---

# Array Characteristics

Important points about arrays in Go:

- Arrays have a fixed size
- All elements must have the same data type
- Array size is part of the array type
- Indexes start at `0`

Example:

```go
var numbers [3]int
```

This array can only store exactly 3 integers.

---

# Declaring Arrays

Syntax:

```go
var arrayName [size]dataType
```

Example:

```go
var numbers [5]int
```

This creates an integer array with 5 elements.

---

# Initializing Arrays

Example:

```go
var numbers [5]int = [5]int{10, 20, 30, 40, 50}
```

Short form:

```go
numbers := [5]int{10, 20, 30, 40, 50}
```

---

# Accessing Array Elements

Use indexes to access elements.

Example:

```go
numbers := [3]int{10, 20, 30}

fmt.Println(numbers[0])
fmt.Println(numbers[1])
```

Output:

```text
10
20
```

---

# Modifying Array Elements

Example:

```go
numbers := [3]int{10, 20, 30}

numbers[1] = 100

fmt.Println(numbers)
```

Output:

```text
[10 100 30]
```

---

# Array Length

Use `len()` to get the array length.

Example:

```go
numbers := [5]int{1, 2, 3, 4, 5}

fmt.Println(len(numbers))
```

Output:

```text
5
```

---

# Automatic Size Detection

Use `...` to let Go count the array size automatically.

Example:

```go
numbers := [...]int{10, 20, 30, 40}
```

Go automatically sets the size to `4`.

---

# Looping Through Arrays

## Using a Traditional `for` Loop

Example:

```go
numbers := [3]int{10, 20, 30}

for i := 0; i < len(numbers); i++ {
    fmt.Println(numbers[i])
}
```

---

# Using `range`

Example:

```go
numbers := [3]int{10, 20, 30}

for index, value := range numbers {
    fmt.Println(index, value)
}
```

Output:

```text
0 10
1 20
2 30
```

---

# Ignoring the Index

Use `_` to ignore unused values.

Example:

```go
for _, value := range numbers {
    fmt.Println(value)
}
```

---

# Multidimensional Arrays

Arrays can contain other arrays.

Example:

```go
var matrix [2][2]int = [2][2]int{
    {1, 2},
    {3, 4},
}
```

Accessing values:

```go
fmt.Println(matrix[0][1])
```

Output:

```text
2
```

---

# Arrays Are Value Types

When assigning one array to another, Go copies all elements.

Example:

```go
a := [3]int{1, 2, 3}
b := a

b[0] = 100

fmt.Println(a)
fmt.Println(b)
```

Output:

```text
[1 2 3]
[100 2 3]
```

The original array does not change.

---

# Example Program

```go
package main

import "fmt"

func main() {
    scores := [5]int{90, 85, 70, 95, 88}

    for _, score := range scores {
        fmt.Println(score)
    }
}
```

---

# Common Beginner Mistakes

## Accessing Invalid Indexes

Incorrect:

```go
numbers := [3]int{1, 2, 3}

fmt.Println(numbers[5])
```

This causes an error because index `5` does not exist.

---

## Wrong Array Size

Incorrect:

```go
numbers := [2]int{1, 2, 3}
```

Too many values for the array size.

---

# Summary

Important points about arrays in Go:

- Arrays store multiple values of the same type
- Arrays have fixed sizes
- Indexes start at `0`
- `len()` returns the array length
- `range` is useful for looping through arrays
- Multidimensional arrays are supported
- Arrays are value types in Go
- Array size is part of the type

Arrays are useful for storing fixed-size collections of data in Golang.