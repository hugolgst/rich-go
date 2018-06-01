package client

import (
	"../ipc"
	"encoding/json"
	"fmt"
)

type Handshake struct {
	V        string `json:"v"`
	ClientId string `json:"client_id"`
}

func Login(clientid string) {
	payload, err := json.Marshal(Handshake{"1", clientid})
	if err != nil {
		panic(err)
	}

	ipc.OpenSocket()
	fmt.Println(ipc.Send(0, string(payload)))
}
