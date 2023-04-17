# GenericSafeMap

GenericSafeMap is a Go library providing a thread-safe map implementation that supports generics, available from Go version 1.18 onwards. This library provides type-safe access to the underlying data structure using generics to ensure that the data stored and retrieved are of the same type.

## Installation

To use this library, you can install it using the go get command:

```
go get github.com/MSSkowron/genericsafemap
```

## Usage

Here's an example of how to use the GenericSafeMap library:

```
package main

import (
    "fmt"
    "github.com/MSSkowron/genericsafemap"
)

func main() {
    m := genericsafemap.New[string, int]()
    m.Put("one", 1)
    m.Put("two", 2)

    val, ok := m.Get("one")
    if !ok {
        fmt.Println("Key not found")
    } else {
        fmt.Println(val)
    }

    m.Remove("one")

    val, ok = m.Get("one")
    if !ok {
        fmt.Println("Key not found")
    } else {
        fmt.Println(val)
    }
}
```

In the above example, we create a new GenericSafeMap instance with string keys and int values, add two key-value pairs to the map, retrieve the value for key "one", remove the key "one", and then attempt to retrieve the value for "one" again.

## License

The library is available as open source under the terms of the MIT License.
