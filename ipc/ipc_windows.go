package ipc

import (
	npipe "gopkg.in/natefinch/npipe.v2"
)

// Open the discord-ipc-0 named pipe
func OpenSocket() error {

	// Connect to the Windows named pipe, this is a well known name
	sock, err := npipe.Dial(`\\.\pipe\discord-ipc-0`)
	if err != nil {
		return err
	}

	socket = sock
	return nil
}
