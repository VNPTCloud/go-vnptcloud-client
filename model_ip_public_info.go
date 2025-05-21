package go-vnptcloud-client

type IpPublicInfo struct {
	Id int32 `json:"id,omitempty"`
	IpAddress string `json:"ipAddress,omitempty"`
	IPv6Address string `json:"iPv6Address,omitempty"`
}
