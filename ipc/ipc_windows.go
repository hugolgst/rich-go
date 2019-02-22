package ipc

import (
	"fmt"

	npipe "gopkg.in/natefinch/npipe.v2"
)

// Open the discord-ipc-0 unix socket
func OpenSocket() error {

	fmt.Println("Windows open socket")

	sock, err := npipe.Dial(`\\.\pipe\discord-ipc-0`)
	//sock, err := npipe.Dial(`\\.\pipe\discord-ipc-0`)
	if err != nil {
		return err
	}

	// sock, err := net.Dial("unix", GetIpcPath()+"/discord-ipc-0")
	// if err != nil {
	// 	return err
	// }

	fmt.Println("Listening to discord-ipc-0")

	socket = sock
	return nil
}
