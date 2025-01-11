package main

import "testing"

func TestPart1(t *testing.T) {
	res := part1("input.txt")
	exp := 388071289
	if res != exp {
		t.Errorf("%d != %d\n", exp, res)
	}
}

func TestPart2(t *testing.T) {
	res := part2("input.txt")
	exp := 84206669
	if res != exp {
		t.Errorf("%d != %d\n", exp, res)
	}
}
