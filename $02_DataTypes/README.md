# Data Types in Golang: A Comprehensive Guide

Go is a **statically typed** and **strongly typed** language. This means every variable has a specific type that is determined at compile time, ensuring memory safety and performance.

---

These are the most fundamental types used for simple values.

### Numeric Types
* **Integers**:
    * Signed: `int8`, `int16`, `int32` (rune), `int64`, and `int` (platform dependent, usually 64-bit).
    * Unsigned: `uint8` (byte), `uint16`, `uint32`, `uint64`, and `uint`.
* **Floating Point**: `float32`, `float64`.
* **Complex**: `complex64`, `complex128`.
* **Special Aliases**:
    * `byte`: Alias for `uint8`, used for raw data.
    * `rune`: Alias for `int32`, represents a Unicode code point/character since Go does not
  have `char` type.

### Boolean Type
* Type: `bool`
* Values: `true` or `false`.

### String Type
* Type: `string`
* Characteristics: Immutable sequence of bytes. Go strings are UTF-8 encoded by default.

---

