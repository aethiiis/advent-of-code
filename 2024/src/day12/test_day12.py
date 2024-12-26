import pytest
from day12 import part1, part2


@pytest.mark.benchmark(group="day12")
def test_part1(benchmark):
    res = benchmark(part1)
    assert res == 1486324


@pytest.mark.benchmark(group="day12")
def test_part2(benchmark):
    res = benchmark(part2)
    assert res == 898684
