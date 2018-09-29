package shopifygraphql

import (
	"encoding/base64"
	"fmt"
)

var shopifyGIDFormat = "gid://shopify/%s/%d"

// ID base64 encode id
func ID(id string) string {
	return base64.URLEncoding.EncodeToString([]byte(id))
}

// ProductID convert product id to graphql id
func ProductID(id int) string {
	return ID(fmt.Sprintf(shopifyGIDFormat, "Product", id))
}

// ProductVariantID convert product variant id to graphql id
func ProductVariantID(id int) string {
	return ID(fmt.Sprintf(shopifyGIDFormat, "ProductVariant", id))
}

// OrderID convert order id to graphql id
func OrderID(id int) string {
	return ID(fmt.Sprintf(shopifyGIDFormat, "Order", id))
}
