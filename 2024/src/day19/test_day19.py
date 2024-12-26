import pytest
from day19 import part1, part2


@pytest.mark.benchmark(group="day19")
def test_part1(benchmark):
    res = benchmark(part1)
    assert res == 300


@pytest.mark.benchmark(group="day19")
def test_part2(benchmark):
    res = benchmark(part2)
    assert res == 624802218898092
