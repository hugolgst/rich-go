package client

import "strconv"

// Error is a Discord error.
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e Error) Error() string {
	return "discord code " + strconv.Itoa(e.Code) + ": " + e.Message
}
