package shopifygraphql

// MetafieldInput metafield input struct
type MetafieldInput struct {
	Description string      `json:"description,omitempty"`
	ID          string      `json:"id,omitempty"`
	Key         string      `json:"key,omitempty"`
	Namespace   string      `json:"namespace,omitempty"`
	Value       interface{} `json:"value,omitempty"`
	ValueType   string      `json:"value_type,omitempty"`
}
