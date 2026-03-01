package errors

import "fmt"

type NIKAlreadyExistsError struct {
	NIK string
}

func (e *NIKAlreadyExistsError) Error() string {
	return fmt.Sprintf("NIK %s sudah terdaftar", e.NIK)
}

func NewNIKAlreadyExistsError(nik string) *NIKAlreadyExistsError {
	return &NIKAlreadyExistsError{NIK: nik}
}
