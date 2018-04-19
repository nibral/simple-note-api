package usecase

type NotPermittedError struct {
	Msg string
}

func (_ *NotPermittedError) Error() string {
	return "Requested operation is not permitted"
}

type InvalidParameterError struct {
	Msg string
}

func (_ *InvalidParameterError) Error() string {
	return "Invalid parameter"
}
