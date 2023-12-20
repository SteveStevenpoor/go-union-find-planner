package planner

import (
	"sort"

	"github.com/SteveStevenpoor/go-union-find-planner/internal/dsu"
)

type Job struct {
	name  string
	dl, s int
}

func Solve(jobs []Job) []Job {
	length := len(jobs) + 1

	sort.Slice(jobs, func(i, j int) bool { return jobs[i].s > jobs[j].s })

	var d dsu.Dsu
	d.DsuInit(length)
	opt := make([]Job, length)
	for r := range jobs {
		dl := jobs[r].dl
		newPlace := d.DsuFind(dl)
		if newPlace != 0 {
			opt[newPlace] = jobs[r]
		} else {
			newPlace = d.DsuFind(length - 1)
			opt[newPlace] = jobs[r]
		}
		d.Parent[newPlace]--
	}

	return opt
}

func GiveTotalStrike(jobs []Job) int {
	strike := 0
	for r := range jobs {
		if jobs[r].dl < r {
			strike += jobs[r].s
		}
	}
	return strike
}

func SolveGreedy(jobs []Job) int {
	sort.Slice(jobs, func(i, j int) bool { return jobs[i].s > jobs[j].s })
	strike := 0
	for r := range jobs {
		if jobs[r].dl < r+1 {
			strike += jobs[r].s
		}
	}
	return strike
}
