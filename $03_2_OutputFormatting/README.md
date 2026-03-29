# Go `fmt` Package Cheat Sheet

This guide provides a comprehensive overview of the formatting "verbs" used in Go's `fmt.Printf`, `fmt.Sprintf`, and `fmt.Fprintf` functions.

---

## 1. General Verbs
These verbs can be used with almost any data type.

| Verb  | Description |
| :---- | :--- |
| `%v`  | **Value**: The default format for the value. |
| `%+v` | **Plus-Value**: For structs, it adds field names. |
| `%#v` | **Go Syntax**: A Go-syntax representation of the value. |
| `%T`  | **Type**: A Go-syntax representation of the type. |
| `%%`  | **Literal**: A literal percent sign; consumes no value. |

## 2. Integer Verbs
Used for whole numbers (int, int64, etc.).

| Verb | Description |
| :--- | :--- |
| `%d` | **Decimal**: Base 10. |
| `%b` | **Binary**: Base 2. |
| `%o` | **Octal**: Base 8. |
| `%O` | **Octal with prefix**: Base 8 with `0o` prefix. |
| `%x` | **Hexadecimal**: Base 16, lowercase (a-f). |
| `%X` | **Hexadecimal**: Base 16, uppercase (A-F). |
| `%U` | **Unicode**: Unicode format `U+1234`. |

## 3. Floating-Point & Complex Verbs
Used for `float32`, `float64`, and complex types.

| Verb | Description |
| :--- | :--- |
| `%f` | **Decimal point**: No exponent (e.g., 123.456). |
| `%e` | **Scientific**: Lowercase `e` (e.g., 1.23456e+78). |
| `%E` | **Scientific**: Uppercase `E` (e.g., 1.23456E+78). |
| `%g` | **General**: Smaller of `%e` or `%f` for better readability. |
| `%G` | **General**: Smaller of `%E` or `%f`. |

## 4. String and Byte Slice Verbs
For text, characters, and memory addresses.

| Verb | Description |
| :--- | :--- |
| `%s` | **String**: Raw bytes of the string or slice. |
| `%q` | **Quoted**: Double-quoted string safely escaped. |
| `%x` | **Hex**: Base 16, lowercase, two chars per byte. |
| `%X` | **Hex**: Base 16, uppercase, two chars per byte. |
| `%c` | **Character**: The Unicode character represented by the integer. |
| `%p` | **Pointer**: 0x-prefixed hexadecimal address. |

## 5. Boolean Verb
| Verb | Description |
| :--- | :--- |
| `%t` | **Boolean**: The word `true` or `false`. |

---

## 6. Width and Precision Control
Modify the behavior of verbs by adding numbers between `%` and the letter.

| Syntax  | Result |
| :------ | :--- |
| `%5d`   | **Width 5**: Right-aligned, padded with spaces. |
| `%-5d`  | **Left-align**: Padded with spaces on the right. |
| `%05d`  | **Zero-padding**: Padded with leading zeros. |
| `%.2f`  | **Precision 2**: Exactly 2 digits after the decimal point. |
| `%9.2f` | **Width 9, Precision 2**: Total width 9, 2 decimal places. |

---

## 7. Quick Code Example
```go
package main

import "fmt"

func main() {
    type User struct {
        ID   int
        Name string
    }

    u := User
    
    // General vs Struct detail
    fmt.Printf("%v\n", u)   // {1 Alice}
    fmt.Printf("%+v\n", u)  // {ID:1 Name:Alice}

    // Number formatting
    fmt.Printf("%03d\n", 7) // 007
    
    // Type checking
    fmt.Printf("%T\n", u)   // main.User
}