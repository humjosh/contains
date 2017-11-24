# Contains

## Example
```go
package main

import (
    "fmt"
    . "github.com/humjosh/contains"
)

func main() {
    vals := []string{
        "Good",
        "morning!",
    }

    if StringSlice(strings).Contains("morning!") {
        fmt.Println("Good morning to you too!")
    }
}
```