package go-vnptcloud-client

type InstanceEditModel struct {
	// Id máy ảo sẽ điều chỉnh
	InstanceId int32 `json:"instanceId,omitempty"`
	// Id hệ điều hành mới
	NewImageId int32 `json:"newImageId,omitempty"`
}
