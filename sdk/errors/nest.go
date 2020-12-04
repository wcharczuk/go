package errors

// Nest nests an arbitrary number of exceptions.
func Nest(err ...error) error {
	var ex *Exception
	var last *Exception
	var didSet bool

	for _, e := range err {
		if e != nil {
			var wrappedEx *Exception
			if typedEx, isTyped := e.(*Exception); !isTyped {
				wrappedEx = &Exception{
					Class:      e,
					StackTrace: Callers(DefaultStartDepth),
				}
			} else {
				wrappedEx = typedEx
			}

			if wrappedEx != ex {
				if ex == nil {
					ex = wrappedEx
					last = wrappedEx
				} else {
					last.Inner = wrappedEx
					last = wrappedEx
				}
				didSet = true
			}
		}
	}
	if didSet {
		return ex
	}
	return nil
}
