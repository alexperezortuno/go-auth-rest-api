package data

type DefaultResponse struct {
	Code    int16    `json:"code,omitempty"`
	Content []string `json:"content,omitempty"`
	Status  string   `json:"status,omitempty"`
}
