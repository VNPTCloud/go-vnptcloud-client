package go-vnptcloud-client

type InstanceRebuildModel struct {
	// Id máy ảo sẽ điều chỉnh
	InstanceId int32 `json:"instanceId,omitempty"`
	// Tên mới của máy ảo
	NewName string `json:"newName,omitempty"`
}
