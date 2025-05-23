package go-vnptcloud-client

type InstanceResizeModel struct {
	// Id máy ảo sẽ điều chỉnh
	InstanceId int32 `json:"instanceId,omitempty"`
	// Cpu tăng thêm
	Cpu int32 `json:"cpu,omitempty"`
	// Ram tăng thêm
	Ram int32 `json:"ram,omitempty"`
	// Dung lượng ổ root tăng thêmn
	VolumeSize int32 `json:"volumeSize,omitempty"`
	// Id loại Gpu tăng thêm
	GpuTypeId int32 `json:"gpuTypeId,omitempty"`
	// Số lượng Gpu tăng thêm
	GpuCount int32 `json:"gpuCount,omitempty"`
}
