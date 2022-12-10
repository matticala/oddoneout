package main

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

const (
	maxN = math.MaxInt8
)

var (
	seed = rand.NewSource(time.Now().UnixMicro())
	r    = rand.New(seed)
)

func generator(reps int) (output []int, unique int) {
	n := r.Intn(maxN)

	output = make([]int, 0, reps*n+1)
	generated := make(map[int]struct{}, n)

	for len(output) < cap(output) {
		n := r.Intn(maxN)
		if _, added := generated[n]; added {
			continue
		}

		if cap(output)-len(output) > 1 {
			for i := 0; i < reps; i++ {
				output = append(output, n)
			}
		} else {
			unique = n
			output = append(output, n)
		}
	}

	rand.Shuffle(len(output), func(i, j int) {
		output[i], output[j] = output[j], output[i]
	})

	return
}

type testFn = func(*testing.T)

func testSolve(reps int) testFn {
	return func(t *testing.T) {
		input, expected := generator(reps)

		n, err := Solve(input, reps)
		if err != nil {
			t.Fatal(err)
		}
		if n != expected {
			t.Log(expected, input)
			t.Fatalf("Solution %d is not what is expected (%d).", n, expected)
		}
	}
}

func TestSolve(t *testing.T) {
	t.Run("Solve with 2 duplicates", testSolve(2))
	t.Run("Solve with 3 duplicates", testSolve(3))
	t.Run("Solve with 4 duplicates", testSolve(4))
}

func BenchmarkSolve(b *testing.B) {
	b.Run("2", func(b *testing.B) {
		input, _ := generator(2)
		if _, err := Solve(input, 2); err != nil {
			b.Fatal(err)
		}
	})
	b.Run("3", func(b *testing.B) {
		input, _ := generator(3)
		if _, err := Solve(input, 3); err != nil {
			b.Fatal(err)
		}
	})
	b.Run("4", func(b *testing.B) {
		input, _ := generator(4)
		if _, err := Solve(input, 4); err != nil {
			b.Fatal(err)
		}
	})

}
