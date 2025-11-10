## 🧾 Print Functions in Go

Go provides three main functions for printing output from the `fmt` package:

| Function      | Adds spaces between args       | Adds newline         | Supports format specifiers |
| ------------- | ------------------------------ | -------------------- | -------------------------- |
| `fmt.Print`   | ❌ No                          | ❌ No                | ❌ No                      |
| `fmt.Println` | ✅ Yes                         | ✅ Yes               | ❌ No                      |
| `fmt.Printf`  | ❌ No (you control via format) | ❌ No (you add `\n`) | ✅ Yes                     |

### Examples

```go
fmt.Print("Hello", "World")       // Output: HelloWorld
fmt.Println("Hello", "World")     // Output: Hello World\n
fmt.Printf("Hello %s\n", "World") // Output: Hello World\n
```
