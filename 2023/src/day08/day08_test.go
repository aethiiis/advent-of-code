package main

import "testing"

func TestPart1(t *testing.T) {
	res := part1("input.txt")
	exp := 20093
	if res != exp {
		t.Errorf("%d != %d\n", exp, res)
	}
}

func TestPart2(t *testing.T) {
	res := part2("input.txt")
	exp := 22103062509257
	if res != exp {
		t.Errorf("%d != %d\n", exp, res)
	}
}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part1("input.txt")
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		part2("input.txt")
	}
}
