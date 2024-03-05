
# Deferrable Module Documentation

## Overview
The `deferrable` module offers utilities to simplify deferred execution of functions in Go, providing a more readable and linear approach to managing deferred calls, especially useful for cleanup and teardown operations.

## Installation
To use the `deferrable` module, add `deferrable` to your project:

```
go get github.com/Patrick-ring-motive/deferrable
```

## Functions

### LinearDefer
Executes a main function immediately and defers the execution of additional functions provided in a slice.

#### Signature
```go
func LinearDefer(mainFunc func(), deferFuncs []func())
```

#### Parameters
- `mainFunc`: The main function to execute immediately.
- `deferFuncs`: A slice of functions to defer.

#### Example Usage
```go
LinearDefer(func() {
    fmt.Println("This should run first.")
}, Defer(func() {
    fmt.Println("This should run last.")
}))
```

### Defer
Wraps deferred functions in a slice to be passed to `LinearDefer`, simplifying the creation of the slice of deferred functions. This serves to show a clear separation between the end of the main function and beginning of the defered functions.

#### Signature
```go
func Defer(deferFuncs ...func()) []func()
```

#### Parameters
- `deferFuncs`: Variadic input of functions to defer.

#### Example Usage
Used in conjunction with `LinearDefer` as shown above.

## Complete Example

```go
package main

import (
    "sync"
    "fmt"
    . "github.com/Patrick-ring-motive/deferrable"
)

func main() {
    var mu sync.Mutex
    LinearDefer(func() {
        fmt.Println("Locking mutex. This should run first.")
        mu.Lock()
    }, Defer(func() {
        fmt.Println("This should run last. Unlocking mutex.")
        mu.Unlock()
    }))
}
```

This example demonstrates using `LinearDefer` to manage mutex locking and unlocking in a readable and linear fashion, improving code clarity and maintainability.
