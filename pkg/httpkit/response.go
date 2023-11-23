package httpkit

type ResponsePayload struct {
	Status   int    `json:"status"`
	ShortUrl string `json:"short_url"`
}
