package models_request

type ResponseModel struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}
