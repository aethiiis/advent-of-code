import pytest
from day13 import part1, part2


@pytest.mark.benchmark(group="day13")
def test_part1(benchmark):
    res = benchmark(part1)
    assert res == 36758


@pytest.mark.benchmark(group="day13")
def test_part2(benchmark):
    res = benchmark(part2)
    assert res == 76358113886726
