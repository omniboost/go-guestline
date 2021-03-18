package guestline

import "fmt"

type ExceptionBlock struct {
	ExceptionCode        int    `xml:"ExceptionCode"`
	ExceptionDescription string `xml:"ExceptionDescription"`
}

func (eb ExceptionBlock) Error() string {
	return fmt.Sprintf("%d: %s", eb.ExceptionCode, eb.ExceptionDescription)
}
