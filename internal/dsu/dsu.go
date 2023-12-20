package dsu

type Dsu struct {
	Parent []int
}

func (d *Dsu) DsuInit(n int) {
	d.Parent = make([]int, n)
	for i := 1; i < n; i++ {
		d.Parent[i] = i
	}
}

func (d *Dsu) DsuFind(x int) int {
	if x != d.Parent[x] {
		d.Parent[x] = d.DsuFind(d.Parent[x])
		return d.Parent[x]
	} else {
		return x
	}
}
