# For Loop in Golang

## Introduction

A loop is used to repeat a block of code multiple times.

In Golang, the main looping structure is the `for` loop.

Unlike some other languages, Go only has the `for` keyword for looping.

---

# Basic `for` Loop Syntax

Syntax:

```go
for initialization; condition; update {
    // code
}
```

Example:

```go
package main

import "fmt"

func main() {
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
    }
}
```

Output:

```text
1
2
3
4
5
```

---

# Parts of a `for` Loop

Example:

```go
for i := 1; i <= 5; i++ {
    fmt.Println(i)
}
```

Explanation:

| Part | Purpose |
|---|---|
| `i := 1` | Initialization |
| `i <= 5` | Condition |
| `i++` | Update |

---

# Infinite Loop

A loop without a condition runs forever.

Example:

```go
for {
    fmt.Println("Running...")
}
```

This is called an infinite loop.

---

# Using `for` Like a `while` Loop

Go does not have a separate `while` keyword.

Example:

```go
i := 1

for i <= 5 {
    fmt.Println(i)
    i++
}
```

Output:

```text
1
2
3
4
5
```

---

# Using `break`

`break` stops the loop immediately.

Example:

```go
for i := 1; i <= 10; i++ {
    if i == 5 {
        break
    }

    fmt.Println(i)
}
```

Output:

```text
1
2
3
4
```

---

# Using `continue`

`continue` skips the current iteration and moves to the next loop cycle.

Example:

```go
for i := 1; i <= 5; i++ {
    if i == 3 {
        continue
    }

    fmt.Println(i)
}
```

Output:

```text
1
2
4
5
```

---

# Nested Loops

A loop inside another loop is called a nested loop.

Example:

```go
for i := 1; i <= 3; i++ {
    for j := 1; j <= 2; j++ {
        fmt.Println(i, j)
    }
}
```

Output:

```text
1 1
1 2
2 1
2 2
3 1
3 2
```

---

# The `range` Keyword

`range` is commonly used to loop through collections like arrays, slices, maps, and strings.

Example with slice:

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

# Ignoring Values with `_`

Use `_` to ignore unused values.

Example:

```go
numbers := []int{10, 20, 30}

for _, value := range numbers {
    fmt.Println(value)
}
```

---

# Example Program

```go
package main

import "fmt"

func main() {
    sum := 0

    for i := 1; i <= 5; i++ {
        sum += i
    }

    fmt.Println("Total:", sum)
}
```

Output:

```text
Total: 15
```

---

# Common Beginner Mistakes

## Forgetting to Update the Loop Variable

Incorrect:

```go
i := 1

for i <= 5 {
    fmt.Println(i)
}
```

This creates an infinite loop.

Correct:

```go
i++
```

---

## Wrong Loop Condition

Incorrect:

```go
for i := 1; i >= 5; i++ {
}
```

The loop will never run.

---

# Summary

Important points about `for` loops in Go:

- Go uses `for` as its main loop structure
- A `for` loop has initialization, condition, and update sections
- Go can simulate a `while` loop using `for`
- `break` exits the loop
- `continue` skips the current iteration
- Nested loops are supported
- `range` is useful for looping through collections
- Infinite loops are possible with `for {}`

`for` loops are essential for repeating tasks and processing data in Golang programs.