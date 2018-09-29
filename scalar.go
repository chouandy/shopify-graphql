package shopifygraphql

import (
	"encoding/base64"

	"github.com/shurcooL/graphql"
)

// ID base64 encode id to GraphQL ID
func ID(id string) graphql.ID {
	return graphql.ID(base64.URLEncoding.EncodeToString([]byte(id)))
}
