package client

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"os"

	"github.com/donovansolms/rich-go-redo/ipc"
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
	Details string `json:"details"`
	State   string `json:"state"`
	Assets  Assets `json:"assets"`
}

type Assets struct {
	LargeImage string `json:"large_image"`
	LargeText  string `json:"large_text"`
	SmallImage string `json:"small_image"`
	SmallText  string `json:"small_text"`
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

		ipc.Send(0, string(payload))
		// fmt.Println(ipc.Send(0, string(payload)))
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
		fmt.Println("Not logged in")
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

	// fmt.Println(string(payload))

	resp := ipc.Send(1, string(payload))
	fmt.Println(resp)
	return nil
}

func getNonce() string {
	buf := make([]byte, 16)
	rand.Read(buf)
	buf[6] = (buf[6] & 0x0f) | 0x40

	return fmt.Sprintf("%x-%x-%x-%x-%x", buf[0:4], buf[4:6], buf[6:8], buf[8:10], buf[10:])
}
