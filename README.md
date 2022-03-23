# Generic validations with Go 1.18

This library provides validation utilities using Go 1.18 generics. This allows you to do type-safe validations.

## Installing

You can install this package using `go mod`:

```
go get go.debugged.it/validation
```

You can then import it:

```go
import "go.debugged.it/validation"
```

## Must

The `Must`, `Must3`, `Must4` functions provide error-to-panic conversion in a type-safe manner:

```go
createWithError := func() (string, error) {
    return "Hello world!", nil
}

fmt.Println(validation.Must(createWithError()))

// Output: "Hello world!"
```

## InSlice

The `InSlice` function searches for elements in slices:

```go
slice := []string{
    "a",
    "b",
    "c",
}
if validation.InSlice(slice, "c") {
    fmt.Println("Yes, c is in the slice")
} else {
    fmt.Println("No, c is not in the slice")
}
```

## NotInSlice

The `NotInSlice` function is the inverse of InSlice above.