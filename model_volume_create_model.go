package go-vnptcloud-client

type VolumeCreateModel struct {
	// Id dự án
	ProjectId int32 `json:"projectId"`
	// Tên máy ảo
	VolumeName string `json:"volumeName"`
	VolumeType string `json:"volumeType"`
	VolumeSize int32 `json:"volumeSize"`
	Description string `json:"description,omitempty"`
	Iops int32 `json:"iops,omitempty"`
	CreateFromSnapshotId int32 `json:"createFromSnapshotId,omitempty"`
	InstanceToAttachId int32 `json:"instanceToAttachId,omitempty"`
	IsMultiAttach bool `json:"isMultiAttach,omitempty"`
	IsEncryption bool `json:"isEncryption,omitempty"`
}
