package go-vnptcloud-client

type PortInfo struct {
	Id string `json:"id,omitempty"`
	FixedIPs []string `json:"fixedIPs,omitempty"`
}
