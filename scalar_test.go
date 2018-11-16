package shopifygraphql

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestID(t *testing.T) {
	// Set test cases
	testCases := []struct {
		gid      string
		expected int
	}{
		{
			gid:      "gid://shopify/Shop/3527475313",
			expected: 3527475313,
		},
		{
			gid:      "gid://shopify/Product/632910392",
			expected: 632910392,
		},
		{
			gid:      "gid://shopify/FulfillmentService/301629553?id=true",
			expected: 301629553,
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.expected, ID(testCase.gid))
		})
	}
}

func TestGID(t *testing.T) {
	// Set test cases
	testCases := []struct {
		resource string
		id       int
		expected string
	}{
		{
			resource: "Shop",
			id:       3527475313,
			expected: "gid://shopify/Shop/3527475313",
		},
		{
			resource: "Product",
			id:       632910392,
			expected: "gid://shopify/Product/632910392",
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.expected, GID(testCase.resource, testCase.id))
		})
	}
}

func TestEncodeGID(t *testing.T) {
	// Set test cases
	testCases := []struct {
		gid      string
		expected string
	}{
		{
			gid:      "gid://shopify/Shop/3527475313",
			expected: "Z2lkOi8vc2hvcGlmeS9TaG9wLzM1Mjc0NzUzMTM=",
		},
		{
			gid:      "gid://shopify/Product/632910392",
			expected: "Z2lkOi8vc2hvcGlmeS9Qcm9kdWN0LzYzMjkxMDM5Mg==",
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.expected, EncodeGID(testCase.gid))
		})
	}
}

func TestEncodedGID(t *testing.T) {
	// Set test cases
	testCases := []struct {
		resource string
		id       int
		expected string
	}{
		{
			resource: "Shop",
			id:       3527475313,
			expected: "Z2lkOi8vc2hvcGlmeS9TaG9wLzM1Mjc0NzUzMTM=",
		},
		{
			resource: "Product",
			id:       632910392,
			expected: "Z2lkOi8vc2hvcGlmeS9Qcm9kdWN0LzYzMjkxMDM5Mg==",
		},
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.expected, EncodedGID(testCase.resource, testCase.id))
		})
	}
}
