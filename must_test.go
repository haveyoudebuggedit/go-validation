package validation_test

import (
	"fmt"

	"github.com/haveyoudebuggedit/validation"
)

func ExampleMust() {
	createWithError := func() (string, error) {
		return "Hello world!", nil
	}

	fmt.Println(validation.Must(createWithError()))

	// Output: Hello world!
}

func ExampleMust3() {
	createWithError := func() (string, string, error) {
		return "Hello ", "world!", nil
	}

	value1, value2 := validation.Must3(createWithError())
	fmt.Printf("%s%s", value1, value2)

	// Output: Hello world!
}

func ExampleMust4() {
	createWithError := func() (string, string, string, error) {
		return "Hello ", "world", "!", nil
	}

	value1, value2, value3 := validation.Must4(createWithError())
	fmt.Printf("%s%s%s", value1, value2, value3)

	// Output: Hello world!
}
