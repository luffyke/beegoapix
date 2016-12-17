package api

//"errors"

type ApiError struct {
	State State
	Err   error
}
