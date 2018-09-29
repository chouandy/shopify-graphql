package shopifygraphql

import (
	"encoding/base64"
	"fmt"

	"github.com/shurcooL/graphql"
)

var shopifyGIDFormat = "gid://shopify/%s/%d"

// ID base64 encode id to GraphQL ID
func ID(id string) graphql.ID {
	return graphql.ID(base64.URLEncoding.EncodeToString([]byte(id)))
}

// ProductID convert product id to graphql id
func ProductID(id int) graphql.ID {
	return ID(fmt.Sprintf(shopifyGIDFormat, "Product", id))
}

// ProductVariantID convert product variant id to graphql id
func ProductVariantID(id int) graphql.ID {
	return ID(fmt.Sprintf(shopifyGIDFormat, "ProductVariant", id))
}

// OrderID convert order id to graphql id
func OrderID(id int) graphql.ID {
	return ID(fmt.Sprintf(shopifyGIDFormat, "Order", id))
}
