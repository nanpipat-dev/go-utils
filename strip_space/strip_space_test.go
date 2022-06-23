package strip_space

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStripSpace(t *testing.T) {

	tests := []struct {
		input  string
		expect string
	}{
		{
			"h e l l o",
			"hello",
		},
		{
			"sund ay",
			"sunday",
		},
		{
			"strip       Space",
			"stripSpace",
		},
	}

	for _, test := range tests {

		got := StripSpace(test.input)

		assert.Equal(t, got, test.expect)
	}
}
