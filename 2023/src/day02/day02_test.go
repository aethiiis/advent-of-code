package main

import "testing"

func TestPart1(t *testing.T) {
	res := part1("input.txt")
	exp := 2239
	if res != exp {
		t.Errorf("%d != %d\n", exp, res)
	}
}

func TestPart2(t *testing.T) {
	res := part2("input.txt")
	exp := 83435
	if res != exp {
		t.Errorf("%d != %d\n", exp, res)
	}
}