# Slices in Golang

## Introduction

A slice is a flexible and dynamic view of an array.

Unlike arrays, slices do not have a fixed size.

Slices are one of the most commonly used data structures in Go.

Example:

```go
numbers := []int{10, 20, 30}
```

---

# Difference Between Arrays and Slices

| Array | Slice |
|---|---|
| Fixed size | Dynamic size |
| Size is part of type | Size is not part of type |
| Less flexible | More flexible |

Example array:

```go
var arr [3]int
```

Example slice:

```go
var slice []int
```

---

# Declaring Slices

Syntax:

```go
var sliceName []dataType
```

Example:

```go
var numbers []int
```

This creates an empty slice.

---

# Initializing Slices

Example:

```go
numbers := []int{10, 20, 30, 40}
```

---

# Accessing Slice Elements

Indexes start at `0`.

Example:

```go
numbers := []int{10, 20, 30}

fmt.Println(numbers[0])
fmt.Println(numbers[1])
```

Output:

```text
10
20
```

---

# Modifying Slice Elements

Example:

```go
numbers := []int{10, 20, 30}

numbers[1] = 100

fmt.Println(numbers)
```

Output:

```text
[10 100 30]
```

---

# Slice Length

Use `len()` to get the number of elements.

Example:

```go
numbers := []int{1, 2, 3, 4}

fmt.Println(len(numbers))
```

Output:

```text
4
```

---

# Slice Capacity

Use `cap()` to get the slice capacity.

Example:

```go
numbers := []int{1, 2, 3}

fmt.Println(cap(numbers))
```

Capacity is the number of elements the slice can hold before resizing.

---

# Creating Slices from Arrays

Syntax:

```go
slice := array[start:end]
```

Example:

```go
arr := [5]int{10, 20, 30, 40, 50}

slice := arr[1:4]

fmt.Println(slice)
```

Output:

```text
[20 30 40]
```

Explanation:

- Start index is included
- End index is excluded

---

# Using `append()`

`append()` adds elements to a slice.

Example:

```go
numbers := []int{1, 2, 3}

numbers = append(numbers, 4)

fmt.Println(numbers)
```

Output:

```text
[1 2 3 4]
```

---

# Appending Multiple Elements

Example:

```go
numbers = append(numbers, 5, 6, 7)
```

---

# Looping Through Slices

## Using `range`

Example:

```go
numbers := []int{10, 20, 30}

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

Example:

```go
for _, value := range numbers {
    fmt.Println(value)
}
```

---

# Copying Slices

Use `copy()` to copy elements.

Example:

```go
source := []int{1, 2, 3}

destination := make([]int, len(source))

copy(destination, source)

fmt.Println(destination)
```

---

# Using `make()`

`make()` creates slices with specific length and capacity.

Syntax:

```go
make([]type, length, capacity)
```

Example:

```go
numbers := make([]int, 3, 5)

fmt.Println(numbers)
```

Output:

```text
[0 0 0]
```

---

# Nil Slices

A slice without initialization is `nil`.

Example:

```go
var numbers []int

fmt.Println(numbers == nil)
```

Output:

```text
true
```

---

# Slices Are Reference Types

Slices reference underlying arrays.

Example:

```go
a := []int{1, 2, 3}
b := a

b[0] = 100

fmt.Println(a)
fmt.Println(b)
```

Output:

```text
[100 2 3]
[100 2 3]
```

Both slices reference the same data.

---

# Example Program

```go
package main

import "fmt"

func main() {
    fruits := []string{"Apple", "Banana", "Orange"}

    fruits = append(fruits, "Mango")

    for _, fruit := range fruits {
        fmt.Println(fruit)
    }
}
```

---

# Common Beginner Mistakes

## Accessing Invalid Indexes

Incorrect:

```go
numbers := []int{1, 2, 3}

fmt.Println(numbers[5])
```

This causes an error.

---

## Forgetting to Store `append()` Result

Incorrect:

```go
append(numbers, 4)
```

Correct:

```go
numbers = append(numbers, 4)
```

---

# Summary

Important points about slices in Go:

- Slices are dynamic collections
- Slices are more flexible than arrays
- `append()` adds new elements
- `len()` returns length
- `cap()` returns capacity
- Slices can be created from arrays
- `range` is commonly used for iteration
- Slices are reference types
- `make()` creates slices with custom size and capacity

Slices are one of the most important and frequently used features in Golang.