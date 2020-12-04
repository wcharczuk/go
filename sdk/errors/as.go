package errors

// As is a helper method that returns an error as an ex.
func As(err interface{}) *Exception {
	if typed, typedOk := err.(Exception); typedOk {
		return &typed
	}
	if typed, typedOk := err.(*Exception); typedOk {
		return typed
	}
	return nil
}
