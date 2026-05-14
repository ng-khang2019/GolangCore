# Maps in Golang

## Introduction

A map in Golang is a collection of key-value pairs.

Each value is stored using a unique key.

Example:

```go
student := map[string]int{
    "Alice": 90,
    "Bob":   85,
}
```

Here:

- Keys → `"Alice"`, `"Bob"`
- Values → `90`, `85`

Maps are useful for fast data lookup.

---

# Why Use Maps?

Maps help:

- Store related data
- Search values quickly
- Associate keys with values
- Organize dynamic data

Maps are similar to dictionaries or hash tables in other languages.

---

# Declaring Maps

Syntax:

```go
var mapName map[keyType]valueType
```

Example:

```go
var ages map[string]int
```

This creates a nil map.

---

# Creating Maps with `make()`

Example:

```go
ages := make(map[string]int)
```

This creates an empty usable map.

---

# Initializing Maps

Example:

```go
scores := map[string]int{
    "Alice": 95,
    "Bob":   88,
}
```

---

# Adding Elements

Example:

```go
scores["John"] = 90
```

---

# Accessing Values

Example:

```go
fmt.Println(scores["Alice"])
```

Output:

```text
95
```

---

# Updating Values

Example:

```go
scores["Alice"] = 100
```

---

# Deleting Elements

Use `delete()`.

Syntax:

```go
delete(map, key)
```

Example:

```go
delete(scores, "Bob")
```

---

# Checking if a Key Exists

Maps return two values:

- The value
- A boolean indicating existence

Example:

```go
value, exists := scores["Alice"]

fmt.Println(value)
fmt.Println(exists)
```

Output:

```text
95
true
```

---

# Map Length

Use `len()` to get the number of elements.

Example:

```go
fmt.Println(len(scores))
```

---

# Looping Through Maps

Use `range`.

Example:

```go
for key, value := range scores {
    fmt.Println(key, value)
}
```

Example output:

```text
Alice 95
Bob 88
```

Important:

Map iteration order is not guaranteed.

---

# Ignoring Keys or Values

Example:

```go
for key := range scores {
    fmt.Println(key)
}
```

Or:

```go
for _, value := range scores {
    fmt.Println(value)
}
```

---

# Nil Maps

A map declared without initialization is nil.

Example:

```go
var data map[string]int

fmt.Println(data == nil)
```

Output:

```text
true
```

Important:

You cannot add elements to a nil map.

Incorrect:

```go
data["A"] = 1
```

This causes a runtime error.

---

# Maps Are Reference Types

Maps share underlying data.

Example:

```go
a := map[string]int{
    "x": 1,
}

b := a

b["x"] = 100

fmt.Println(a)
fmt.Println(b)
```

Output:

```text
map[x:100]
map[x:100]
```

Both variables reference the same map.

---

# Nested Maps

Maps can contain other maps.

Example:

```go
students := map[string]map[string]int{
    "Alice": {
        "Math": 90,
        "English": 85,
    },
}
```

Accessing values:

```go
fmt.Println(students["Alice"]["Math"])
```

---

# Example Program

```go
package main

import "fmt"

func main() {
    capitals := map[string]string{
        "Japan": "Tokyo",
        "France": "Paris",
    }

    capitals["Vietnam"] = "Hanoi"

    for country, capital := range capitals {
        fmt.Println(country, ":", capital)
    }
}
```

---

# Common Beginner Mistakes

## Using Nil Maps

Incorrect:

```go
var data map[string]int

data["A"] = 1
```

Correct:

```go
data := make(map[string]int)
```

---

## Expecting Ordered Output

Maps do not maintain insertion order.

Incorrect expectation:

```text
Alice
Bob
Charlie
```

The order may change each run.

---

# Summary

Important points about maps in Go:

- Maps store key-value pairs
- Keys must be unique
- `make()` creates usable maps
- Values are accessed using keys
- `delete()` removes elements
- `range` loops through maps
- Map iteration order is not guaranteed
- Maps are reference types
- Nil maps cannot store values

Maps are essential for efficient data storage and lookup in Golang applications.