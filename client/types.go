package client

import (
	"encoding/json"
)

const (
	handshakeOpcode uint32 = iota
	frameOpcode
	closeOpcode
)

type handshake struct {
	V        string `json:"v"`
	ClientID string `json:"client_id"`
}

type frame struct {
	Cmd   string          `json:"cmd"`
	Args  args            `json:"args"`
	Data  json.RawMessage `json:"data"`
	Evt   string          `json:"evt"`
	Nonce string          `json:"nonce"`
}

type args struct {
	Pid      int              `json:"pid"`
	Activity *payloadActivity `json:"activity"`
}

type payloadActivity struct {
	Details    string             `json:"details,omitempty"`
	State      string             `json:"state,omitempty"`
	Assets     payloadAssets      `json:"assets,omitempty"`
	Party      *payloadParty      `json:"party,omitempty"`
	Timestamps *payloadTimestamps `json:"timestamps,omitempty"`
	Secrets    *payloadSecrets    `json:"secrets,omitempty"`
	Buttons    []*payloadButton   `json:"buttons,omitempty"`
}

type payloadAssets struct {
	LargeImage string `json:"large_image,omitempty"`
	LargeText  string `json:"large_text,omitempty"`
	SmallImage string `json:"small_image,omitempty"`
	SmallText  string `json:"small_text,omitempty"`
}

type payloadParty struct {
	ID   string `json:"id,omitempty"`
	Size [2]int `json:"size,omitempty"`
}

type payloadTimestamps struct {
	Start *uint64 `json:"start,omitempty"`
	End   *uint64 `json:"end,omitempty"`
}

type payloadSecrets struct {
	Match    string `json:"match,omitempty"`
	Join     string `json:"join,omitempty"`
	Spectate string `json:"spectate,omitempty"`
}

type payloadButton struct {
	Label string `json:"label,omitempty"`
	URL   string `json:"url,omitempty"`
}
