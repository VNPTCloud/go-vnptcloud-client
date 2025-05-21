package go-vnptcloud-client

type VlanInfo struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	SubnetAddressRequired string `json:"subnetAddressRequired,omitempty"`
}
