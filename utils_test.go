package segtree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_expand(t *testing.T) {
	testCases := []struct {
		desc                       string
		arr                        []int
		to                         int
		expectedArr, expectedExtra []int
	}{
		{
			desc:          "len(arr) < to",
			arr:           []int{1, 2, 3, 4},
			to:            6,
			expectedArr:   []int{1, 2, 3, 4},
			expectedExtra: []int{0, 0},
		},
		{
			desc:          "len(arr) == to",
			arr:           []int{1, 2, 3},
			to:            3,
			expectedArr:   []int{1, 2, 3},
			expectedExtra: nil,
		},
		{
			desc:          "len(arr) > to",
			arr:           []int{1, 2, 3, 4, 5},
			to:            2,
			expectedArr:   []int{1, 2, 3, 4, 5},
			expectedExtra: nil,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actualArr, actualExtra := expand(tC.arr, tC.to)
			require.Equal(t, tC.expectedArr, actualArr)
			require.Equal(t, tC.expectedExtra, actualExtra)
		})
	}
}

func Test_nextPowerOf2(t *testing.T) {
	testCases := []struct {
		desc     string
		num      int
		expected int
	}{
		{
			desc:     "already power of 2",
			num:      32,
			expected: 32,
		},
		{
			desc:     "3, should return 4",
			num:      3,
			expected: 4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			require.Equal(t, tC.expected, nextPowerOf2(tC.num))
		})
	}
}
