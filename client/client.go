package client

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"os"

	"github.com/heroslender/rich-go/ipc"
)

var isLoggedIn bool

func Login(clientid string) error {
	if isLoggedIn == false {
		payload, err := json.Marshal(Handshake{"1", clientid})
		if err != nil {
			return err
		}

		err = ipc.OpenSocket()
		if err != nil {
			return err
		}

		// TODO: Response should be parsed
		ipc.Send(0, string(payload))
	}
	isLoggedIn = true

	return nil
}

func Logout() {
	isLoggedIn = false
	ipc.CloseSocket()
}

func SetActivity(activity Activity) error {
	if isLoggedIn == false {
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

	// TODO: Response should be parsed
	ipc.Send(1, string(payload))
	return nil
}

func getNonce() string {
	buf := make([]byte, 16)
	rand.Read(buf)
	buf[6] = (buf[6] & 0x0f) | 0x40

	return fmt.Sprintf("%x-%x-%x-%x-%x", buf[0:4], buf[4:6], buf[6:8], buf[8:10], buf[10:])
}
