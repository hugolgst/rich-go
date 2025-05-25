package client

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"os"

	"github.com/hugolgst/rich-go/ipc"
)

var logged bool

// Login sends a handshake in the socket and returns an error or nil
func Login(clientid string) error {
	if !logged {
		payload, err := json.Marshal(Handshake{"1", clientid})
		if err != nil {
			return err
		}

		err = ipc.OpenSocket()
		if err != nil {
			return err
		}

		if err := sendWithError(0, string(payload)); err != nil {
			return err
		}
	}
	logged = true

	return nil
}

func Logout() {
	logged = false

	err := ipc.CloseSocket()
	if err != nil {
		panic(err)
	}
}

func SetActivity(activity Activity) error {
	if !logged {
		return nil
	}

	payload, err := json.Marshal(Frame{
		"SET_ACTIVITY",
		Args{
			os.Getpid(),
			mapActivity(&activity),
		},
		getNonce(),
	})

	if err != nil {
		return err
	}

	if err := sendWithError(1, string(payload)); err != nil {
		return err
	}
	return nil
}

func getNonce() string {
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	if err != nil {
		fmt.Println(err)
	}

	buf[6] = (buf[6] & 0x0f) | 0x40

	return fmt.Sprintf("%x-%x-%x-%x-%x", buf[0:4], buf[4:6], buf[6:8], buf[8:10], buf[10:])
}

func sendWithError(opcode int, payload string) error {
	resp := ipc.Send(opcode, payload)

	var data struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	if err := json.Unmarshal([]byte(resp), &data); err != nil {
		return err
	}

	if data.Code == 0 {
		return nil
	}

	return fmt.Errorf("error code %d: %s", data.Code, data.Message)
}
