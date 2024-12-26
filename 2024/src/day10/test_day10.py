import pytest
from day10 import part1, part2


@pytest.mark.benchmark(group="day10")
def test_part1(benchmark):
    res = benchmark(part1)
    assert res == 652


@pytest.mark.benchmark(group="day10")
def test_part2(benchmark):
    res = benchmark(part2)
    assert res == 1432
