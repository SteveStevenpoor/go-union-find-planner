package planner

import (
	"fmt"
	"testing"
)

func TestSolve(t *testing.T) {
	jobA := Job{
		name: "A",
		dl:   3,
		s:    25,
	}

	jobB := Job{
		name: "B",
		dl:   4,
		s:    10,
	}

	jobC := Job{
		name: "C",
		dl:   1,
		s:    30,
	}

	jobD := Job{
		name: "D",
		dl:   3,
		s:    50,
	}

	jobE := Job{
		name: "E",
		dl:   3,
		s:    20,
	}

	// TEST1
	jobs := []Job{jobA, jobB, jobC, jobD, jobE}
	opt := Solve(jobs)
	actual := GiveTotalStrike(opt)
	expected := 20

	assert(t, actual, expected)
	fmt.Printf("TEST1 ||||| Uf: %d, Greedy: %d\n", actual, SolveGreedy(jobs))

	// TEST2
	jobA.dl = 2
	jobA.s = 10

	jobB.dl = 1
	jobB.s = 5

	jobC.dl = 3
	jobC.s = 8

	jobD.dl = 2
	jobD.s = 3

	jobE.dl = 1
	jobE.s = 2

	jobs = []Job{jobA, jobB, jobC, jobD, jobE}
	opt = Solve(jobs)
	actual = GiveTotalStrike(opt)
	expected = 5

	assert(t, actual, expected)
	fmt.Printf("TEST2 ||||| Uf: %d, Greedy: %d\n", actual, SolveGreedy(jobs))

	// TEST3
	jobA.dl = 4
	jobA.s = 15

	jobB.dl = 1
	jobB.s = 20

	jobC.dl = 2
	jobC.s = 70

	jobD.dl = 2
	jobD.s = 60

	jobE.dl = 2
	jobE.s = 50

	jobs = []Job{jobA, jobB, jobC, jobD, jobE}
	opt = Solve(jobs)
	actual = GiveTotalStrike(opt)
	expected = 70

	assert(t, actual, expected)
	fmt.Printf("TEST3 ||||| Uf: %d, Greedy: %d\n", actual, SolveGreedy(jobs))

	// TEST4
	jobA.dl = 2
	jobA.s = 19

	jobB.dl = 1
	jobB.s = 7

	jobC.dl = 3
	jobC.s = 13

	jobD.dl = 5
	jobD.s = 25

	jobE.dl = 1
	jobE.s = 41

	jobs = []Job{jobA, jobB, jobC, jobD, jobE}
	opt = Solve(jobs)
	actual = GiveTotalStrike(opt)
	expected = 7

	assert(t, actual, expected)
	fmt.Printf("TEST4 ||||| Uf: %d, Greedy: %d\n", actual, SolveGreedy(jobs))

	// TEST5
	jobA.dl = 3
	jobA.s = 2

	jobB.dl = 1
	jobB.s = 7

	jobC.dl = 4
	jobC.s = 39

	jobD.dl = 1
	jobD.s = 12

	jobE.dl = 1
	jobE.s = 55

	jobs = []Job{jobA, jobB, jobC, jobD, jobE}
	opt = Solve(jobs)
	actual = GiveTotalStrike(opt)
	expected = 19

	assert(t, actual, expected)
	fmt.Printf("TEST5 ||||| Uf: %d, Greedy: %d\n", actual, SolveGreedy(jobs))
}

func assert(t *testing.T, actual, expected any) {
	if expected != actual {
		t.Errorf("Result was incorrect, got: %d, want: %d.", actual, expected)
	}
}
