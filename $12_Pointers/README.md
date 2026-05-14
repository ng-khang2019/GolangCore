# Pointers in Golang

## Introduction

A pointer is a variable that stores the memory address of another variable.

Instead of storing a value directly, a pointer points to where the value is stored in memory.

Pointers are useful for:

- Modifying variables directly
- Improving performance
- Sharing data between functions
- Working with structs and large data

---

# Memory Address

Every variable is stored somewhere in memory.

Example:

```go
x := 10
```

`x` has:

- A value → `10`
- A memory address

---

# The `&` Operator

The `&` operator gets the memory address of a variable.

Example:

```go
package main

import "fmt"

func main() {
    x := 10

    fmt.Println(x)
    fmt.Println(&x)
}
```

Possible output:

```text
10
0xc0000120a0
```

The hexadecimal value is the memory address.

---

# Declaring Pointers

Syntax:

```go
var pointerName *dataType
```

Example:

```go
var ptr *int
```

This pointer can store the address of an integer.

---

# Creating a Pointer

Example:

```go
x := 10

ptr := &x
```

Here:

- `x` stores `10`
- `ptr` stores the address of `x`

---

# Dereferencing Pointers

The `*` operator accesses the value stored at a memory address.

Example:

```go
package main

import "fmt"

func main() {
    x := 10

    ptr := &x

    fmt.Println(*ptr)
}
```

Output:

```text
10
```

---

# Modifying Values Through Pointers

Pointers can change the original variable value.

Example:

```go
package main

import "fmt"

func main() {
    x := 10

    ptr := &x

    *ptr = 50

    fmt.Println(x)
}
```

Output:

```text
50
```

Changing `*ptr` also changes `x`.

---

# Pointer Example with Functions

Without pointers:

```go
func changeValue(num int) {
    num = 100
}
```

The original value does not change.

---

# Using Pointers in Functions

Example:

```go
package main

import "fmt"

func changeValue(num *int) {
    *num = 100
}

func main() {
    x := 10

    changeValue(&x)

    fmt.Println(x)
}
```

Output:

```text
100
```

Explanation:

- `&x` passes the memory address
- `*num` changes the original value

---

# Nil Pointers

A pointer without an assigned address is `nil`.

Example:

```go
var ptr *int

fmt.Println(ptr)
```

Output:

```text
<nil>
```

---

# Checking for Nil

Example:

```go
if ptr != nil {
    fmt.Println(*ptr)
}
```

This prevents runtime errors.

---

# Pointer to Struct

Pointers are commonly used with structs.

Example:

```go
type Person struct {
    Name string
}

func main() {
    person := Person{Name: "Alice"}

    ptr := &person

    fmt.Println(ptr.Name)
}
```

Go automatically dereferences struct pointers.

---

# Using `new()`

`new()` allocates memory and returns a pointer.

Example:

```go
ptr := new(int)

fmt.Println(*ptr)
```

Output:

```text
0
```

The default value for `int` is `0`.

---

# Pointer Comparison

Pointers can be compared.

Example:

```go
a := 10
b := 10

ptr1 := &a
ptr2 := &b

fmt.Println(ptr1 == ptr2)
```

Output:

```text
false
```

Different variables have different memory addresses.

---

# Example Program

```go
package main

import "fmt"

func increase(num *int) {
    *num++
}

func main() {
    value := 5

    increase(&value)

    fmt.Println(value)
}
```

Output:

```text
6
```

---

# Common Beginner Mistakes

## Dereferencing Nil Pointers

Incorrect:

```go
var ptr *int

fmt.Println(*ptr)
```

This causes a runtime error.

---

## Forgetting `&`

Incorrect:

```go
changeValue(x)
```

Correct:

```go
changeValue(&x)
```

---

## Confusing `*` Usage

`*` has two meanings:

| Usage | Meaning |
|---|---|
| `*int` | Pointer type |
| `*ptr` | Dereference pointer |

---

# Summary

Important points about pointers in Go:

- Pointers store memory addresses
- `&` gets a variable's address
- `*` dereferences a pointer
- Pointers can modify original variables
- Pointers are useful in functions and structs
- Nil pointers do not point to valid memory
- `new()` creates pointers with allocated memory
- Go simplifies pointer usage compared to some other languages

Pointers are an important concept for writing efficient and flexible Golang programs.