package segtree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewTree(t *testing.T) {
	arr := []int{-2, 0, 3, -5, 2, -1}
	tree := NewTree(arr, func(a, b int) int { return a + b }, 0, nil)
	require.Equal(t, tree.arr, []int{-3, -4, 1, -2, -2, 1, 0, -2, 0, 3, -5, 2, -1, 0, 0})
	require.Equal(t, tree.rng, Range{l: 0, r: 8})
}
