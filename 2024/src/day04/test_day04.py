import pytest
from day04 import part1, part2


@pytest.mark.benchmark(group="day04")
def test_part1(benchmark):
    res = benchmark(part1)
    assert res == 2454


@pytest.mark.benchmark(group="day04")
def test_part2(benchmark):
    res = benchmark(part2)
    assert res == 1858
