package ipc

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

var socket net.Conn

func CloseSocket() error {
	if socket != nil {
		socket.Close()
		socket = nil
	}
	return nil
}

// Read the socket response
func Read() (string, error) {
	buf := make([]byte, 512)
	payloadlength, err := socket.Read(buf)
	if err != nil {
		//fmt.Println("Nothing to read")
		return "", err
	}

	buffer := new(bytes.Buffer)
	buffer.Write(buf[8:payloadlength])

	return buffer.String(), nil
}

// Send opcode and payload to the unix socket
func Send(opcode int, payload string) error {
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
	_, err = socket.Write(buf.Bytes())
	if err != nil {
		return err
	}

	return nil
}
