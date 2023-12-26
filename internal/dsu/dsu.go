package dsu

type Dsu struct {
	Parent, Rank, Left []int
}

func (d *Dsu) DsuInit(n int) {
	d.Parent = make([]int, n)
	d.Left = make([]int, n)
	for i := 1; i < n; i++ {
		d.Parent[i] = i
		d.Left[i] = i
	}
	d.Rank = make([]int, n)
}

func (d *Dsu) DsuUnion(x, y int) {
	parentX, parentY := d.DsuFind(x), d.DsuFind(y)

	if d.Rank[parentX] > d.Rank[parentY] {
		d.Parent[parentY] = parentX
	} else {
		d.Parent[parentX] = parentY
		if d.Rank[parentX] == d.Rank[parentY] {
			d.Rank[parentY] += 1
		}
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
