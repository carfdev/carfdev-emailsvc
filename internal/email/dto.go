package email

type SendResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
