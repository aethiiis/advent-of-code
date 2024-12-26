import pytest
from day09 import part1, part2


@pytest.mark.benchmark(group="day09")
def test_part1(benchmark):
    res = benchmark(part1)
    assert res == 6332189866718


@pytest.mark.benchmark(group="day09")
def test_part2(benchmark):
    res = benchmark(part2)
    assert res == 6353648390778
