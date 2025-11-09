//go:build !windows
// +build !windows

package ipc

import (
	"fmt"
	"net"
	"time"
)

// Dials discord-pic-[0,9] sockets and returns an error if none connect.
func OpenSocket() error {
	var err error
	for i := 0; i < 10; i++ {
		if sock, err := net.DialTimeout("unix", GetIpcPath()+fmt.Sprintf("/discord-ipc-%d", i), time.Second*2); err == nil {
			socket = sock
			return nil
		}
	}
	return fmt.Errorf("dialed sockets discord-ipc-[0,9], all failed, last error %w", err)
}
