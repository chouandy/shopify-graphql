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
	}

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("TestCase[%d]", i+1), func(t *testing.T) {
			assert.Equal(t, testCase.expected, ID(testCase.gid))
		})
	}
}
