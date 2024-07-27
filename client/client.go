package client

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/dikey0ficial/rich-go/ipc"
	"os"
)

var logged bool

// Login sends a handshake in socket and returns an error or nil
func Login(clientid string) error {
	var gerr error
	if !logged {
		payload, err := json.Marshal(Handshake{"1", clientid})
		if err != nil {
			return err
		}

		err = ipc.OpenSocket()
		if err != nil {
			return err
		}

		// TODO: Response should be parsed
		_, gerr = ipc.Send(0, string(payload))
	}
	logged = true

	return gerr
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

	// TODO: Response should be parsed
	_, err = ipc.Send(1, string(payload))
	return err
}

func getNonce() (string, error) {
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	if err != nil {
		fmt.Println(err)
	}

	buf[6] = (buf[6] & 0x0f) | 0x40

	return fmt.Sprintf("%x-%x-%x-%x-%x", buf[0:4], buf[4:6], buf[6:8], buf[8:10], buf[10:])
}
