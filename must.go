package validation

// Must takes two arguments: a value and an error. If the error is not nil, it panics with the error, otherwise it
// returns the value.
//
// Example usage:
//
//     result := validation.Must(functionThatMayReturnAnError())
func Must[K any](value K, err error) K {
	if err != nil {
		panic(err)
	}
	return value
}

// Must3 takes three arguments: two values and an error. If the error is not nil, it panics with the error, otherwise it
// returns the two values.
//
// Example usage:
//
//     result1, result2 := validation.Must3(functionThatMayReturnAnError())
func Must3[K any, L any](value1 K, value2 L, err error) (K, L) {
	if err != nil {
		panic(err)
	}
	return value1, value2
}

// Must4 takes four arguments: three values and an error. If the error is not nil, it panics with the error, otherwise
// it returns the three values.
//
// Example usage:
//
//     result1, result2, result3 := validation.Must4(functionThatMayReturnAnError())
func Must4[K any, L any, M any](value1 K, value2 L, value3 M, err error) (K, L, M) {
	if err != nil {
		panic(err)
	}
	return value1, value2, value3
}
