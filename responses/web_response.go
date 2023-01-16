package responses

type WebResponse struct {
	Code    interface{} `json:"code,omitempty"`
	Status  interface{} `json:"status,omitempty"`
	Message interface{} `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
