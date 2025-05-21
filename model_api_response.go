
package go-vnptcloud-client

type ModelApiResponse struct {
	Success bool `json:"success,omitempty"`
	Code int32 `json:"code,omitempty"`
	Data *Object `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
	ErrorCode string `json:"errorCode,omitempty"`
}
