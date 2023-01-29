package core

type HTTPErrorResponse struct {
	Message *string `json:"message,omitempty"`
	Result  string  `json:"result"`
}

type HTTPDefaultResponse struct {
	Result string `json:"result"`
}
