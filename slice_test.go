package validation_test

import (
	"fmt"

	"go.debugged.it/validation"
)

func ExampleInSlice() {
	slice := []string{
		"a",
		"b",
		"c",
	}
	if validation.InSlice("c", slice) {
		fmt.Println("Yes, c is in the slice")
	} else {
		fmt.Println("No, c is not in the slice")
	}

	// Output: Yes, c is in the slice
}

func ExampleNotInSlice() {
	slice := []string{
		"a",
		"b",
		"c",
	}
	if validation.NotInSlice("d", slice) {
		fmt.Println("No, d is not in the slice")
	} else {
		fmt.Println("Yes, d is in the slice")
	}

	// Output: No, d is not in the slice
}
