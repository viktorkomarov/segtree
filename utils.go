package segtree

func expand[n any](arr []n, to int) ([]n, []n) {
	if to-len(arr) <= 0 {
		return arr, nil
	}
	extra := make([]n, to-len(arr))
	return arr, extra
}

func nextPowerOf2(num int) int {
	k := 1
	for k < num {
		k = k << 1
	}
	return k
}
