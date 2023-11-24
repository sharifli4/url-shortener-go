package httpkit

type UrlPayload struct {
	Url string `json:"url" validate:"required,url"`
}
