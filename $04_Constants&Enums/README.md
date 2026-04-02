# Constants and Enums in Golang

---
## Constants
### Definition
**Constants** are fixed values that cannot be changed after declaration. They must be known at **compile-time**.

### Characteristics
**Constants** are declared with the `const` keyword. They are fixed values that the program may not alter during its execution.
* Constants must be assigned a value at declaration.

* Values must be compile-time constants (numbers, strings, booleans).

* Cannot use the `:=` syntax (which is only for variables).
* Untyped constants will be implicitly converted to the type of the value they represent.
```go
package main

import "fmt"

// Global Constants (Visible to the whole package)
const DefaultPort = 8080
const AppName = "GopherEngine"

func main() {
// Local Constants (Scoped to this function)
const Version = "1.0.5"

    fmt.Printf("Running %s v%s on port %d\n", AppName, Version, DefaultPort)
}
```
## The `iota` identifier
`iota` is a special built-in constant that is incremented by one each time it is used.
```go
const (
// Stars with a default value of 0
    Guest = iota // 0
    User         // 1
    Admin        // 2
)

const (
    _ = iota // Skip zero
    Frist // 1
    Second // 2
    Third // 3
)

// Bitwise shift with iota
const (
    Read = 1 << iota  // 1 (0001)
    Write             // 2 (0010)
    Execute           // 4 (0100)
)

```
## Enums
### Definition
Since Go does not have a built-in `enum` keyword, we create enums using custom types and `iota`.
### Enum Pattern
By defining a new type based on `int`, we can ensure type safety for our enum values.
```go
type OrderStatus int

const (
    Pending OrderStatus = iota // 0
    Processed                    // 1
    Shipped                      // 2
    Delivered                    // 3
)
```
## Why use `iota`?
* **Auto-increment**: It starts at 0 and increments by 1 for each line in a const block.

* **Skip values**: Use the blank identifier _ to skip a value.

* **Custom start**: Start from 1 by using iota + 1.
```go
type Day int

const (
	/* The first line can be declared as '_ = iota' without a
	 * type. Its type will be inferred as "untyped integer".
	 * Therefore, the following constants will be of type int.
	 */
    _ Day = iota          // Skip 0
    Monday                // 1
    Tuesday               // 2
    Wednesday             // 3
    Thursday              // 4
    Friday                // 5
    Saturday              // 6
    Sunday                // 7
)
```