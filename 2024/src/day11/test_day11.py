import pytest
from day11 import part1, part2


@pytest.mark.benchmark(group="day11")
def test_part1(benchmark):
    res = benchmark(part1)
    assert res == 233050


@pytest.mark.benchmark(group="day11")
def test_part2(benchmark):
    res = benchmark(part2)
    assert res == 276661131175807
