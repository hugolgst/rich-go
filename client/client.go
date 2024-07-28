package client

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"os"

	"github.com/EpicStep/discord-game-sdk-go/transport"
)

var (
	conn   transport.Conn
	logged bool
)

// Login sends a handshake in the socket and returns an error or nil.
func Login(clientId string) error {
	if logged {
		return nil
	}

	ctx := context.Background()

	payload, err := json.Marshal(Handshake{"1", clientId})
	if err != nil {
		return err
	}

	conn, err = transport.New(ctx, transport.Options{})
	if err != nil {
		return fmt.Errorf("failed to create connection: %w", err)
	}

	if err = conn.Write(ctx, handshakeOpcode, payload); err != nil {
		return fmt.Errorf("failed to send handshake: %w", err)
	}

	opcode, data, err := conn.Read(ctx)
	if err != nil {
		return fmt.Errorf("failed to read data from conn: %w", err)
	}

	// when we received non frame - is an error.
	if opcode != frameOpcode {
		if opcode == closeOpcode {
			var closeError Error

			if err = json.Unmarshal(data, &closeError); err != nil {
				return err
			}

			return closeError
		}

		return fmt.Errorf("unexpected opcode %d", opcode)
	}

	logged = true

	return nil
}

func Logout() error {
	if !logged {
		return nil
	}

	logged = false

	err := conn.Close()
	if err != nil {
		return fmt.Errorf("failed to close connection: %w", err)
	}

	return nil
}

func SetActivity(activity Activity) error {
	if !logged {
		return nil
	}

	ctx := context.Background()

	payload, err := json.Marshal(Frame{
		Cmd: "SET_ACTIVITY",
		Args: Args{
			os.Getpid(),
			mapActivity(&activity),
		},
		Nonce: getNonce(),
	})

	if err != nil {
		return fmt.Errorf("failed to marshal activity: %w", err)
	}

	if err = conn.Write(ctx, frameOpcode, payload); err != nil {
		return fmt.Errorf("failed to write activity: %w", err)
	}

	opcode, rawResp, err := conn.Read(ctx)
	if err != nil {
		return fmt.Errorf("failed to read data from conn: %w", err)
	}

	if opcode != frameOpcode {
		return fmt.Errorf("unexpected opcode %d", opcode)
	}

	var resp Frame

	if err = json.Unmarshal(rawResp, &resp); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if resp.Evt == "ERROR" {
		var e Error

		if err = json.Unmarshal(resp.Data, &e); err != nil {
			return err
		}

		return e
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
