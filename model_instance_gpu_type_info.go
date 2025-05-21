package go-vnptcloud-client

// Loại GPU
type InstanceGpuTypeInfo struct {
	// Id loại GPU
	Id int32 `json:"id,omitempty"`
	// Tên loại GPU
	Name string `json:"name"`
}
