// +build !windows

package ipc

import (
	"net"
	"os"
	"time"
)

// Choose the right directory to the ipc socket and return it
func getIpcPath() string {
	variablesnames := []string{"XDG_RUNTIME_DIR", "TMPDIR", "TMP", "TEMP"}

	for _, variablename := range variablesnames {
		path, exists := os.LookupEnv(variablename)

		if exists {
			return path
		}
	}

	return "/tmp"
}

// OpenSocket opens the discord-ipc-0 unix socket
func OpenSocket() error {
	sock, err := net.DialTimeout("unix", getIpcPath()+"/discord-ipc-0", time.Second*2)
	if err != nil {
		return err
	}

	socket = sock
	return nil
}
