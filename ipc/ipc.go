package ipc

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
)

var socket net.Conn

// Choose the right directory to the ipc socket and return it
func GetIpcPath() string {
	variablesnames := []string{"XDG_RUNTIME_DIR", "TMPDIR", "TMP", "TEMP"}

	for _, variablename := range variablesnames {
		path, exists := os.LookupEnv(variablename)

		if exists {
			return path
		}
	}

	return "/tmp"
}

// Open the discord-ipc-0 unix socket
func OpenSocket() {
	sock, err := net.Dial("unix", GetIpcPath()+"/discord-ipc-0")
	if err != nil {
		panic(err)
	}

	fmt.Println("Listening to discord-ipc-0")

	socket = sock
}

// Read the socket response
func Read() string {
	buf := make([]byte, 512)
	payloadlength, err := socket.Read(buf)
	if err != nil {
		fmt.Println("Nothing to read")
	}

	buffer := new(bytes.Buffer)
	for i := 8; i < payloadlength; i++ {
		buffer.WriteByte(buf[i])
	}

	return string(buffer.Bytes())
}

// Send opcode and payload to the unix socket
func Send(opcode int, payload string) string {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, int32(opcode))
	if err != nil {
		fmt.Println(err)
	}

	err = binary.Write(buf, binary.LittleEndian, int32(len(payload)))
	if err != nil {
		fmt.Println(err)
	}

	buf.Write([]byte(payload))
	socket.Write(buf.Bytes())

	return Read()
}
