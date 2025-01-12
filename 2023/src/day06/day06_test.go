package main

import "testing"

func TestPart1(t *testing.T) {
	res := part1("input.txt")
	exp := 588588
	if res != exp {
		t.Errorf("%d != %d\n", exp, res)
	}
}

func TestPart2(t *testing.T) {
	res := part2("input.txt")
	exp := 34655848
	if res != exp {
		t.Errorf("%d != %d\n", exp, res)
	}
}
