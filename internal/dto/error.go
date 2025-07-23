package dto

type ErrorResponse struct {
	Message string            `json:"message" example:"request error"`
	Fields  map[string]string `json:"fields,omitempty"`
}
