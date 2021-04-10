package ipc

import (
	"bytes"
	"encoding/binary"
	"net"
	"os"
)

// IPC is a Inter Process Communications client
type IPC struct {
	socket net.Conn
}

// NewIPC returns a IPC client
func NewIPC() (*IPC, error) {
	ipc := IPC{}

	if err := ipc.openSocket(); err != nil {
		return nil, err
	}

	return &ipc, nil
}

// Choose the right directory to the ipc socket and return it
func (ipc *IPC) getIpcPath() string {
	variablesnames := []string{"XDG_RUNTIME_DIR", "TMPDIR", "TMP", "TEMP"}

	for _, variablename := range variablesnames {
		path, exists := os.LookupEnv(variablename)

		if exists {
			return path
		}
	}

	return "/tmp"
}

// CloseSocket close IPC socket
func (ipc *IPC) CloseSocket() error {
	if ipc.socket != nil {
		if err := ipc.socket.Close(); err != nil {
			return err
		}

		ipc.socket = nil
	}
	return nil
}

// Read the socket response
func (ipc *IPC) Read() []byte {
	buf := make([]byte, 512)
	payloadlength, _ := ipc.socket.Read(buf)

	return buf[8:payloadlength]
}

// Send opcode and payload to the unix socket
func (ipc *IPC) Send(opcode int, payload string) ([]byte, error) {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, int32(opcode))
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.LittleEndian, int32(len(payload)))
	if err != nil {
		return nil, err
	}

	buf.Write([]byte(payload))
	_, err = ipc.socket.Write(buf.Bytes())
	if err != nil {
		return nil, err
	}

	return ipc.Read(), nil
}
