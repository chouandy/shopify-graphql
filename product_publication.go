package shopifygraphql

import "time"

// ProductPublicationInput product publication input struct
type ProductPublicationInput struct {
	PublicationID string     `json:"publicationId,omitempty"`
	PublishDate   *time.Time `json:"publishDate,omitempty"`
}
