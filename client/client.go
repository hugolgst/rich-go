package client

import (
	"github.com/ananagame/rich-go/ipc"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"os"
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
	State   string `json:"state"`
	Details string `json:"details"`
}

func Login(clientid string) {
	payload, err := json.Marshal(Handshake{"1", clientid})
	if err != nil {
		panic(err)
	}

	ipc.OpenSocket()
	fmt.Println(ipc.Send(0, string(payload)))
}

func SetActivity(activity Activity) {
	payload, err := json.Marshal(Frame{
		"SET_ACTIVITY",
		Args{
			os.Getpid(),
			activity,
		},
		getNonce(),
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(string(payload))

	fmt.Println(ipc.Send(1, string(payload)))
}


func getNonce() string {
	buf := make([]byte, 16)
	rand.Read(buf)
	buf[6] = (buf[6] & 0x0f) | 0x40

	return fmt.Sprintf("%x-%x-%x-%x-%x", buf[0:4], buf[4:6], buf[6:8], buf[8:10], buf[10:])
}