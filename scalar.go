package shopifygraphql

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

var shopifyGIDFormat = "gid://shopify/%s/%d"

// EncodeGID base64 encode id
func EncodeGID(gid string) string {
	return base64.URLEncoding.EncodeToString([]byte(gid))
}

// ID retrieve id
func ID(gid string) int {
	// Split gid with /
	substrings := strings.Split(gid, "/")
	// Convert str id to int
	id, _ := strconv.Atoi(substrings[len(substrings)-1])

	return id
}

// ShopID convert shop id to graphql id
func ShopID(id int) string {
	return EncodeGID(fmt.Sprintf(shopifyGIDFormat, "Shop", id))
}

// ProductID convert product id to graphql id
func ProductID(id int) string {
	return EncodeGID(fmt.Sprintf(shopifyGIDFormat, "Product", id))
}

// ProductVariantID convert product variant id to graphql id
func ProductVariantID(id int) string {
	return EncodeGID(fmt.Sprintf(shopifyGIDFormat, "ProductVariant", id))
}

// OrderID convert order id to graphql id
func OrderID(id int) string {
	return EncodeGID(fmt.Sprintf(shopifyGIDFormat, "Order", id))
}
