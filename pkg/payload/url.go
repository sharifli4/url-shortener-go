package payload

type UrlPayload struct {
	Url string `json:"url" validate:"required"`
}
