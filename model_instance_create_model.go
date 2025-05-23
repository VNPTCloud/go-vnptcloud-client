package go-vnptcloud-client

type InstanceCreateModel struct {
	// Số lượng máy ảo
	InstanceCount int32 `json:"instanceCount,omitempty"`
	// Tên máy ảo
	InstanceName string `json:"instanceName"`
	// Mật khẩu máy ảo
	InstancePassword string `json:"instancePassword,omitempty"`
	KeyPairId int32 `json:"keyPairId,omitempty"`
	// Id hệ điều hành
	ImageId int32 `json:"imageId,omitempty"`
	// Cpu
	Cpu int32 `json:"cpu,omitempty"`
	// Ram
	Ram int32 `json:"ram,omitempty"`
	// Dung lượng ổ root
	VolumeSize int32 `json:"volumeSize,omitempty"`
	// ssd hay hdd
	VolumeType string `json:"volumeType"`
	// Mã hóa volume
	VolumeEncryption bool `json:"volumeEncryption,omitempty"`
	// Id loại Gpu
	GpuTypeId int32 `json:"gpuTypeId,omitempty"`
	// Số lượng Gpu
	GpuCount int32 `json:"gpuCount,omitempty"`
	// Danh sách SG gắn vào máy ảo
	SecurityGroups []string `json:"securityGroups,omitempty"`
	// Ip Public đã mua
	IpPublicId int32 `json:"ipPublicId,omitempty"`
	// VLAN Id
	VlanId string `json:"vlanId"`
	// Port của VLAN
	PortId string `json:"portId,omitempty"`
	SnapshotId int32 `json:"snapshotId,omitempty"`
}
