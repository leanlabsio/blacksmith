package model

const (
	BUILD_STATUS_RUNNING uint32 = iota
	BUILD_STATUS_SUCCESS
	BUILD_STATUS_FAILED
)

type Build struct {
	Commit   string `json:"commit"`
	UserName string `json:"username"`
	Log      string `json:"log,omitempty"`
	Status   uint32 `json:"state"`
}
