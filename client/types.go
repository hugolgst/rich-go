package client

// Handshake is a login request to Discord
type Handshake struct {
	V        string `json:"v"`
	ClientID string `json:"client_id"`
}

// Frame is a packet to Discord
type Frame struct {
	Cmd   string `json:"cmd"`
	Args  Args   `json:"args"`
	Nonce string `json:"nonce"`
}

// Args is a arguments of Frame
type Args struct {
	Pid      int              `json:"pid"`
	Activity *PayloadActivity `json:"activity"`
}

// PayloadActivity is a change user activity request
type PayloadActivity struct {
	Details    string             `json:"details"`
	State      string             `json:"state"`
	Assets     PayloadAssets      `json:"assets"`
	Party      *PayloadParty      `json:"party,omitempty"`
	Timestamps *PayloadTimestamps `json:"timestamps,omitempty"`
	Secrets    *PayloadSecrets    `json:"secrets,omitempty"`
	Buttons    []*PayloadButton   `json:"buttons,omitempty"`
}

// PayloadAssets is a images in PayloadActivity
type PayloadAssets struct {
	LargeImage string `json:"large_image"`
	LargeText  string `json:"large_text"`
	SmallImage string `json:"small_image"`
	SmallText  string `json:"small_text"`
}

// PayloadParty is a discord party config in PayloadActivity
type PayloadParty struct {
	ID   string `json:"id"`
	Size [2]int `json:"size"`
}

// PayloadTimestamps is a timestamps field in PayloadActivity
type PayloadTimestamps struct {
	Start *uint64 `json:"start,omitempty"`
	End   *uint64 `json:"end,omitempty"`
}

// PayloadSecrets is a secrets in PayloadActivity
type PayloadSecrets struct {
	Match    string `json:"match,omitempty"`
	Join     string `json:"join,omitempty"`
	Spectate string `json:"spectate,omitempty"`
}

// PayloadButton is a button config in PayloadActivity
type PayloadButton struct {
	Label string `json:"label,omitempty"`
	URL   string `json:"url,omitempty"`
}
