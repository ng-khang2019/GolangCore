# Input and Output in Golang

---
## Output
`fmt` package provides some functions to output data to the console:
* `Print()`/`Println()`
* `Printf()`, `Sprintlnf()` and `Sprintf()`
Basics verbs for printing values:
* | Verb  | Description |
  | :---- | :--- |
  | `%v`  | **Value**: The default format for the value. |
  | `%+v` | **Plus-Value**: For structs, it adds field names. |
  | `%#v` | **Go Syntax**: A Go-syntax representation of the value. |
  | `%T`  | **Type**: A Go-syntax representation of the type. |
  | `%%`  | **Literal**: A literal percent sign; consumes no value. |
Look out for more in the "Output Formatting" folder
## Input

* `fmt.Scan(&variable)`: Scans the next space-separated value into the variable.
* `fmt.Scanln(&variable)`: Similar to Scan() but stops at the end of the line.
* `fmt.Scanf()`: Scans with a very specific format.