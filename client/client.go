package client

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"os"

	"github.com/heroslender/rich-go/ipc"
)

type Handshake struct {
	V        string `json:"v"`
	ClientId string `json:"client_id"`
}

type Frame struct {
	Cmd   string `json:"cmd"`
	Args  Args   `json:"args"`
	Nonce string `json:"nonce"`
}

type Args struct {
	Pid      int      `json:"pid"`
	Activity Activity `json:"activity"`
}

type Activity struct {
	Details    string     `json:"details"`
	State      string     `json:"state"`
	Assets     Assets     `json:"assets,omitempty"`
	Party      Party      `json:"party,omitempty"`
	Timestamps Timestamps `json:"timestamps,omitempty"`
	Secrets    Secrets    `json:"secrets,omitempty"`
}

type Assets struct {
	LargeImage string `json:"large_image"`
	LargeText  string `json:"large_text"`
	SmallImage string `json:"small_image"`
	SmallText  string `json:"small_text"`
}

type Party struct {
	ID   string `json:"id"`
	Size [2]int `json:"size"`
}

type Timestamps struct {
	Start int64 `json:"start,omitempty"`
	End   int64 `json:"end,omitempty"`
}

type Secrets struct {
	Match    string `json:"match,omitempty"`
	Join     string `json:"join,omitempty"`
	Spectate string `json:"spectate,omitempty"`
}

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
			activity,
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
