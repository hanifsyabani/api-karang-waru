package errors

import "fmt"

type NIKAlreadyExistsError struct {
	NIK string
}

func (e *NIKAlreadyExistsError) Error() string {
	return fmt.Sprintf("NIK %s sudah digunakan", e.NIK)
}
func NewNIKAlreadyExistsError(nik string) error {
	return &NIKAlreadyExistsError{NIK: nik}
}
