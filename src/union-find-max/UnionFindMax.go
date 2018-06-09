package unionfindmax

// UnionFindMax type
type UnionFindMax struct {
	rootArr []int
	sizeArr []int
	maxArr  []int
}

// NewUnionFindMax instance
func NewUnionFindMax(n int) *UnionFindMax {
	newRootArr := make([]int, n, n)
	newMaxArr := make([]int, n, n)
	newSizeArr := make([]int, n, n)
	for i := 0; i < n; i++ {
		newRootArr[i] = i
		newMaxArr[i] = i
		newSizeArr[i] = 1
	}
	return &UnionFindMax{newRootArr, newSizeArr, newMaxArr}
}

func (qu *UnionFindMax) getRoot(p int) int {
	i := p
	for qu.rootArr[i] != i {
		qu.rootArr[i] = qu.rootArr[qu.rootArr[i]]
		i = qu.rootArr[i]
	}
	return i
}

// Find provides the maximum element within the tree.
func (qu *UnionFindMax) Find(p int) int {
	rootP := qu.getRoot(p)
	return qu.maxArr[rootP]
}

// Union places the small size tree under the larger one.
func (qu *UnionFindMax) Union(p int, q int) {
	rootP := qu.getRoot(p)
	rootQ := qu.getRoot(q)

	maxPQ := func() int {
		if qu.maxArr[rootP] > qu.maxArr[rootQ] {
			return qu.maxArr[rootP]
		}
		return qu.maxArr[rootQ]
	}()

	if qu.sizeArr[p] < qu.sizeArr[q] {
		qu.rootArr[rootP] = rootQ
		qu.sizeArr[rootQ] = qu.sizeArr[rootQ] + qu.sizeArr[rootP]
		qu.maxArr[rootQ] = maxPQ
	} else {
		qu.rootArr[rootQ] = rootP
		qu.sizeArr[rootP] = qu.sizeArr[rootP] + qu.sizeArr[rootQ]
		qu.maxArr[rootP] = maxPQ

	}
}

func (qu *UnionFindMax) isConnected(p int, q int) bool {
	return qu.getRoot(p) == qu.getRoot(q)
}
