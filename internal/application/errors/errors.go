package custom_errors

var (
	OrderNotFoundError        = &orderNotFoundError{}
	OrderAlreadyCanceledError = &orderAlreadyCanceledError{}
	OrderAlreadyFilledError   = &orderAlreadyFilledError{}
)

type orderNotFoundError struct {
	cause error
}

func NewOrderNotFoundError(err error) *orderNotFoundError {
	return &orderNotFoundError{
		cause: err,
	}

}

func (e orderNotFoundError) Error() string {
	return "order record not found"
}

type orderAlreadyCanceledError struct {
	cause error
}

func NewOrderAlreadyCanceledError(err error) *orderAlreadyCanceledError {
	return &orderAlreadyCanceledError{
		cause: err,
	}
}

func (e orderAlreadyCanceledError) Error() string {
	return "order already canceled"
}

type orderAlreadyFilledError struct {
	cause error
}

func NewOrderAlreadyFilledError(err error) *orderAlreadyFilledError {
	return &orderAlreadyFilledError{
		cause: err,
	}
}

func (e orderAlreadyFilledError) Error() string {
	return "order already filled"
}
