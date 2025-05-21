package go-vnptcloud-client

type ImageInfo struct {
	Name string `json:"name,omitempty"`
	Versions []ImageVersionInfo `json:"versions,omitempty"`
}
