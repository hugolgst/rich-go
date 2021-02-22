package client

import "fmt"

type ErrorCode int

func (e ErrorCode) Error() string {
	return fmt.Sprintf("rich-go: error code %d", e)
}

// Discord error codes.
const (
	NoErr ErrorCode = 0
)

type Error struct {
	Code    ErrorCode    `json:"code"`
	Message string `json:"message"`

	Data struct {
		Code    ErrorCode    `json:"code"`
		Message string `json:"message"`
	} `json:"data"`
}

func (e Error) Error() string {
	if e.Code != NoErr {
		return fmt.Sprintf("rich-go: %s", e.Message)
	}

	return fmt.Sprintf("rich-go: %s", e.Data.Message)
}

func (e Error) getCode() ErrorCode {
	if e.Code != NoErr {
		return e.Code
	}

	return e.Data.Code
}