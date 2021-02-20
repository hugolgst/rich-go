package client

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"os"

	"github.com/hugolgst/rich-go/ipc"
)

type Client struct {
	ipc *ipc.IPC
	logged bool
}

func NewClient(clientID string) (*Client, error) {
	c := Client{}

	i, err := ipc.NewIPC()
	if err != nil {
		return nil, err
	}

	c.ipc = i

	if err := c.login(clientID); err != nil {
		return nil, err
	}

	return &c, nil
}

// Login sends a handshake in the socket and returns an error or nil
func (c *Client) login(clientid string) error {
	if !c.logged {
		payload, err := json.Marshal(Handshake{"1", clientid})
		if err != nil {
			return err
		}

		// TODO: Response should be parsed
		c.ipc.Send(0, string(payload))
	}
	c.logged = true

	return nil
}

func (c *Client) Logout() {
	c.logged = false

	err := c.ipc.CloseSocket()
	if err != nil {
		panic(err)
	}
}

func (c *Client) SetActivity(activity Activity) error {
	if !c.logged {
		return nil
	}

	payload, err := json.Marshal(Frame{
		"SET_ACTIVITY",
		Args{
			os.Getpid(),
			mapActivity(&activity),
		},
		c.getNonce(),
	})

	if err != nil {
		return err
	}

	// TODO: Response should be parsed
	c.ipc.Send(1, string(payload))
	return nil
}

func (c *Client) getNonce() string {
	buf := make([]byte, 16)
	_, err := rand.Read(buf)
	if err != nil {
		fmt.Println(err)
	}

	buf[6] = (buf[6] & 0x0f) | 0x40

	return fmt.Sprintf("%x-%x-%x-%x-%x", buf[0:4], buf[4:6], buf[6:8], buf[8:10], buf[10:])
}

