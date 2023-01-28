package core

type SASCreditAuthResponse struct {
	Token string  `json:"token"`
	Error *string `json:"error"`
}

type StatusForCheckRequest struct {
	Token string `json:"token" formdata:"token"`
}

type StatusForCheckResponse struct {
	StatusList []string `json:"status_list"`
	Error      *string  `json:"error"`
}

type CalculateDebtRequest struct {
	Token  string `json:"token" formdata:"token"`
	Number string `json:"number" formdata:"number"`
}

type CalculateDebtResponse struct {
	Success bool    `json:"success"`
	Error   *string `json:"error"`
}
