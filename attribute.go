package shopifygraphql

import "github.com/shurcooL/graphql"

// AttributeInput attribute input
type AttributeInput struct {
	Key   graphql.String `json:"key,omitempty"`
	Value graphql.String `json:"value,omitempty"`
}
