import pytest
from day01 import part1, part2


@pytest.mark.benchmark(group="day01")
def test_part1(benchmark):
    res = benchmark(part1)
    assert res == 1590491


@pytest.mark.benchmark(group="day01")
def test_part2(benchmark):
    res = benchmark(part2)
    assert res == 22588371
