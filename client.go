package shopifygraphql

import (
	"net/http"

	"github.com/shurcooL/graphql"
)

/* ------------------------- */
/*           Admin           */
/* ------------------------- */

// AdminAccessTokenHeader admin access token header
var AdminAccessTokenHeader = "X-Shopify-Access-Token"

// AdminTransport admin transport
type AdminTransport struct {
	AccessToken string
}

// RoundTrip http Transport RoundTrip interface
func (t *AdminTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set(AdminAccessTokenHeader, t.AccessToken)
	return http.DefaultTransport.RoundTrip(req)
}

// NewAdminClient new admin graphql client
func NewAdminClient(domain, accessToken string) *graphql.Client {
	// New http client with transport
	httpClient := &http.Client{
		Transport: &AdminTransport{
			AccessToken: accessToken,
		},
	}
	// Build graphql url
	url := "https://" + domain + "/admin/api/graphql.json"

	// New graphql client
	return graphql.NewClient(url, httpClient)
}

/* ------------------------------ */
/*           Storefront           */
/* ------------------------------ */

// StorefrontAccessTokenHeader storefront access token header
var StorefrontAccessTokenHeader = "X-Shopify-Storefront-Access-Token"

// StorefrontTransport storefront transport
type StorefrontTransport struct {
	AccessToken string
}

// RoundTrip http Transport RoundTrip interface
func (t *StorefrontTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set(StorefrontAccessTokenHeader, t.AccessToken)
	return http.DefaultTransport.RoundTrip(req)
}

// NewStorefrontClient new storefront graphql client
func NewStorefrontClient(domain, accessToken string) *graphql.Client {
	// New http client with transport
	httpClient := &http.Client{
		Transport: &StorefrontTransport{
			AccessToken: accessToken,
		},
	}
	// Build graphql url
	url := "https://" + domain + "/api/graphql.json"

	// New graphql client
	return graphql.NewClient(url, httpClient)
}
