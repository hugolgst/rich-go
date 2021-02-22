package client

import (
	"crypto/rand"
	"encoding/json"
	"errors"
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

	if err := c.Login(clientID); err != nil {
		return nil, err
	}

	return &c, nil
}

// Login sends a handshake in the socket and returns an error or nil
func (c *Client) Login(clientid string) error {
	if !c.logged {
		if err := c.handler(0, Handshake{"1", clientid}); err != nil {
			return err
		}
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
		return errors.New("rich-go: client not login")
	}

	payload := Frame{
		"SET_ACTIVITY",
		Args{
			os.Getpid(),
			mapActivity(&activity),
		},
		c.getNonce(),
	}

	if err := c.handler(1, payload); err != nil {
		return err
	}

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

func (c *Client) handler(opcode int, payload interface{}) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	var response Error

	resp, err := c.ipc.Send(opcode, string(jsonPayload))
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(resp), &response); err != nil {
		return err
	}

	switch response.getCode() {
	case NoErr:
		return nil
	}

	return &response
}

