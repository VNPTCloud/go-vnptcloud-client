package go-vnptcloud-client

type VolumeResizeModel struct {
	VolumeId int32 `json:"volumeId"`
	NewSize int32 `json:"newSize"`
	Iops int32 `json:"iops"`
	Description string `json:"description,omitempty"`
	ProjectId int32 `json:"projectId"`
	VolumeName string `json:"volumeName,omitempty"`
}
