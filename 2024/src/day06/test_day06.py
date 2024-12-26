import pytest
from day06 import part1, part2


@pytest.mark.benchmark(group="day06")
def test_part1(benchmark):
    res = benchmark(part1)
    assert res == 5162


@pytest.mark.benchmark(group="day06")
def test_part2(benchmark):
    res = benchmark(part2)
    assert res == 1909
