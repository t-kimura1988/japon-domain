package exceptions

import "fmt"

type DataNotFound struct {
	Message string
	Code    int
}

func (err *DataNotFound) Error() string {
	return fmt.Sprintf("%s", err.Message)
}
