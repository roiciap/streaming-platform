package http_model

type CreditentialsRequest struct {
	Login    string `json:"login" validate:"min=1"`
	Password string `json:"password" validate:"min=1"`
}
