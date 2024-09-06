package exceptions

type BindingError struct {
	Err error
}

func (e BindingError) Error() string {
	return e.Err.Error()
}
