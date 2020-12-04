/*
Package errors provides the foundations for error handling in the SDK tree.

To create an error that includes a given string class and stack trace:

	err := errors.New("this is a structured error")
	...
	fmt.Println(errors.ErrStackTrace(err))


When in doubt, wrap any errors from non-sdk methods with an exception:

	res, err := http.Get(...)
	if err != nil {
		return nil, errors.New(err) // stack trace will originate from this ca..
	}

To create an error from a known error class, that can be used later to check the type of the error:

	var ErrTooManyFoos errors.Class = "too many foos"
	...
	err := errors.New(ErrTooManyFoos)
	...
	if errors.Is(err, ErrTooManyFoos) { // we can now verify the type of the err with `errors.Is(err, class)`
		fmt.Println("We did too many foos!")
	}

We can pass other options to the `errors.New(...)` constructor, such as setting an inner error:

	err := errors.New(ErrValidation, errors.OptInner(err))
	...
	if errors.Is(err, ErrValidation) {
		fmt.Printf("validation error: %v\n", errors.ErrInner(err))
	}
*/
package errors // import "go.charczuk.com/sdk/errors"
