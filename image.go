package shopifygraphql

// ImageInput image input struct
type ImageInput struct {
	AltText string `json:"altText,omitempty"`
	ID      string `json:"id,omitempty"`
	Src     string `json:"src,omitempty"`
}
