// +build !windows

package ipc

import (
	"net"
	"time"
)

// OpenSocket opens the discord-ipc-0 unix socket
func (ipc *IPC) openSocket() error {
	sock, err := net.DialTimeout("unix", ipc.GetIpcPath()+"/discord-ipc-0", time.Second*2)
	if err != nil {
		return err
	}

	ipc.socket = sock
	return nil
}
