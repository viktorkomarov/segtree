package segtree

type ConnectiveFunc[n any] func(a, b n) n

type ChangeRangeFunc[n any] func(stored n, val n) n

type Tree[n any] struct {
	arr         []n
	propagation []n
	updateFn    ChangeRangeFunc[n]
	connective  ConnectiveFunc[n]
	neutral     n
	rng         Range
}

type Range struct {
	l, r int // [l, r)
}

func (r Range) mid() int {
	return r.l + (r.r-r.l)/2
}

func leftChild(idx int) int {
	return 2*idx + 1
}

func rightChild(idx int) int {
	return 2*idx + 2
}

func parent(idx int) int {
	return (idx - 1) / 2
}

func NewTree[n any](arr []n, connective ConnectiveFunc[n], neutral n, updateFn ChangeRangeFunc[n]) Tree[n] {
	powerOf2 := nextPowerOf2(len(arr))
	arr, tail := expand(arr, powerOf2)
	for i := range tail {
		tail[i] = neutral
	}
	arr, head := expand(append(arr, tail...), 2*powerOf2-1)
	arr = append(head, arr...)

	for i := len(arr) - 1; i > 0; i -= 2 {
		arr[parent(i)] = connective(arr[i], arr[i-1])
	}
	propagation := make([]n, len(arr))
	for i := range propagation {
		propagation[i] = neutral
	}

	return Tree[n]{
		arr:         arr,
		propagation: propagation,
		updateFn:    updateFn,
		connective:  connective,
		neutral:     neutral,
		rng:         Range{l: 0, r: powerOf2},
	}
}

func (t Tree[n]) Range(l, r int) n {
	return t.rangeBorders(0, t.rng, Range{l: l, r: r})
}

func (t Tree[n]) rangeBorders(treeIdx int, treeRange, searchRange Range) n {
	if treeRange.r <= searchRange.l || treeRange.l >= searchRange.r {
		return t.neutral
	}

	if searchRange.l <= treeRange.l && treeRange.r <= searchRange.r {
		return t.arr[treeIdx]
	}

	mid := treeRange.mid()
	return t.connective(
		t.rangeBorders(leftChild(treeIdx), Range{l: treeRange.l, r: mid}, searchRange),
		t.rangeBorders(rightChild(treeIdx), Range{l: mid, r: treeRange.r}, searchRange),
	)
}

func (t Tree[n]) Update(idx int, val n) {
	t.update(idx, val, 0, t.rng)
}

func (t Tree[n]) update(idx int, val n, treeIdx int, treeRange Range) n {
	if idx < treeRange.l || idx >= treeRange.r {
		return t.arr[treeIdx]
	}

	if treeRange.l+1 == treeRange.r {
		t.arr[treeIdx] = t.updateFn(t.arr[treeIdx], val)
		return t.arr[treeIdx]
	}

	mid := treeRange.mid()
	t.arr[treeIdx] = t.connective(
		t.update(idx, val, leftChild(treeIdx), Range{l: treeRange.l, r: mid}),
		t.update(idx, val, rightChild(treeIdx), Range{l: mid, r: treeRange.r}),
	)
	return t.arr[treeIdx]
}

func (t Tree[n]) UpdateRangeLazy(rng Range, val n) {
	t.updateRange(rng, val, 0, t.rng)
}

func (t Tree[n]) updateRange(rng Range, val n, treeIdx int, treeRange Range) {
	if treeRange.r <= rng.l || treeRange.l >= rng.r {
		return
	}

	if rng.l <= treeRange.l && treeRange.r <= rng.r {
		t.propagation[treeIdx] = t.updateFn(t.propagation[treeIdx], val)
		return
	}

	t.propagate(treeIdx, treeRange)
	mid := treeRange.mid()
	t.updateRange(rng, val, leftChild(treeIdx), Range{l: treeRange.l, r: mid})
	t.updateRange(rng, val, rightChild(treeIdx), Range{l: mid, r: treeRange.r})
}

func (t Tree[n]) propagate(treeIdx int, treeRange Range) {
	t.propagation[leftChild(treeIdx)] = t.updateFn(t.propagation[leftChild(treeIdx)], t.propagation[treeIdx])
	t.propagation[rightChild(treeIdx)] = t.updateFn(t.propagation[rightChild(treeIdx)], t.propagation[treeIdx])
	t.propagation[treeIdx] = t.neutral
}
