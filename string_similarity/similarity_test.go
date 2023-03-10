package string_similarity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Similarity(t *testing.T) {
	testCases := []struct {
		name string
		str1 string
		str2 string
		want float32
		have float32
	}{
		{
			name: "should be 100% of coincidence",
			str1: "hello, world!",
			str2: "hello, world!",
			want: 100,
		},
		{
			name: "should be 100% of coincidence",
			str1: "Run go mod tidy, which removes any dependencies the module might have accumulated that are no longer necessary.",
			str2: "Run go mod tidy, which removes any dependencies the module might have accumulated that are no longer necessary.",
			want: 100,
		},
	}
	for _, tc := range testCases {
		assert.Equal(t, tc.want, Similarity(tc.str1, tc.str2))
	}
}
