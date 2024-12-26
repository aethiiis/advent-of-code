import pytest
from day21 import part1, part2


@pytest.mark.benchmark(group="day21")
def test_part1(benchmark):
    res = benchmark(part1)
    assert res == 176452


@pytest.mark.benchmark(group="day21")
def test_part2(benchmark):
    res = benchmark(part2)
    assert res == 218309335714068
