package ipc

import (
	"net"
)

// Open the discord-ipc-0 unix socket
func OpenSocket() error {

	sock, err := net.Dial("unix", GetIpcPath()+"/discord-ipc-0")
	if err != nil {
		return err
	}

	socket = sock
	return nil
}
