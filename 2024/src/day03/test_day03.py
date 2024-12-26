import pytest
from day03 import part1, part2


@pytest.mark.benchmark(group="day03")
def test_part1(benchmark):
    res = benchmark(part1)
    assert res == 168539636


@pytest.mark.benchmark(group="day03")
def test_part2(benchmark):
    res = benchmark(part2)
    assert res == 97529391
