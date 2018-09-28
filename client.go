package shopifygraphql

import (
	"fmt"
	"net/http"

	"github.com/shurcooL/graphql"
)

// NewClient new graphql client
func NewClient(domain, storefrontAccessToken string) *graphql.Client {
	// New http client with transport
	httpClient := &http.Client{
		Transport: &Transport{
			StorefrontAccessToken: storefrontAccessToken,
		},
	}
	// Build graphql url
	url := fmt.Sprintf("https://%s/api/graphql", domain)

	// New graphql client
	return graphql.NewClient(url, httpClient)
}
