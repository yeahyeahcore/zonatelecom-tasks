package core

type HTTPErrorResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}
