package shopifygraphql

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

/* ---------------------- */
/*           ID           */
/* ---------------------- */

var shopifyGIDFormat = "gid://shopify/%s/%d"

// ID retrieve id
func ID(gid string) int {
	// Split gid with /
	substrings := strings.Split(gid, "/")
	// Convert str id to int
	id, _ := strconv.Atoi(substrings[len(substrings)-1])

	return id
}

// GID retrieve id
func GID(resource string, id int) string {
	return fmt.Sprintf(shopifyGIDFormat, resource, id)
}

// EncodeGID base64 encode gid
func EncodeGID(gid string) string {
	return base64.URLEncoding.EncodeToString([]byte(gid))
}

// EncodedGID base64 encoded gid
func EncodedGID(resource string, id int) string {
	return EncodeGID(GID(resource, id))
}

/* ------------------------ */
/*           Type           */
/* ------------------------ */

// Bool bool
func Bool(v bool) *bool { return &v }

// BoolValue bool value
func BoolValue(v *bool) bool { return *v }
