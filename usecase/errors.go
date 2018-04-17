package usecase

type NotPermittedError struct {
	Msg string
}

func (_ *NotPermittedError) Error() string {
	return "Requested operation is not permitted"
}
