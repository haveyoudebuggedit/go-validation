# Generic validations with Go 1.18

This library provides validation utilities using Go 1.18 generics. This allows you to do type-safe validations.

## Installing

You can install this package using `go mod`:

```
go get go.debugged.it/validation
```

You can then import it:

```
import go.debugged.it/validation
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
