package js

type JsError struct {
	Err error
}

func (e JsError) Error() string {
	return e.Err.Error()
}

func Throw(err error) {
	panic(JsError{Err: err})
}
