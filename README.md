# Contains
Helper types `StringSlice`, `IntSlice`, `Int32Slice` and `Int64Slice` each have a `Contains` method, allowing you to easily check if a value is contained within a slice.

## Examples
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

    if StringSlice(vals).Contains("morning!") {
        fmt.Println("Good morning to you too!")
    }

    nums := []int64{
	1, 2, 3, 4, 5,
    }

    if Int64Slice(nums).Contains(6) {
	panic("Shouldn't see this!")
    }
}
```
