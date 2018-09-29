package shopifygraphql

import (
	"net/http"

	"github.com/chouandy/goex/fmtex"
)

// StorefrontAccessTokenHeader storefront access token header
var StorefrontAccessTokenHeader = "X-Shopify-Storefront-Access-Token"

// Transport token transport
type Transport struct {
	StorefrontAccessToken string
}

// RoundTrip http Transport RoundTrip interface
func (t *Transport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add(StorefrontAccessTokenHeader, t.StorefrontAccessToken)
	fmtex.StructPrintln(req)
	return http.DefaultTransport.RoundTrip(req)
}
