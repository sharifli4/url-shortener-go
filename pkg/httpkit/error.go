package httpkit

type ErrorPayload struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
