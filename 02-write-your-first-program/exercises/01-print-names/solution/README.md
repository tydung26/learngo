## 🧾 Print Functions in Go

Go provides three main functions for printing output from the `fmt` package (fun fact: `fmt` is short for **“format”**, because it provides Go’s built-in formatting tools for input and output):

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

## 💡 Fun Fact

The `fmt` package is part of Go’s standard library, and its name comes from the word **format**, because its main job is **formatting input and output** — printing things to the console or reading formatted data.
Think of it like this:

- `f` = format
- `mt` = methods / tools

So `fmt` ≈ **formatting tools**.
It powers everything from simple `Println()` calls to advanced formatted strings like `Printf()` and `Sprintf()`.
