import pytest
from day25 import part1, part2


@pytest.mark.benchmark(group="day25")
def test_part1(benchmark):
    res = benchmark(part1)
    assert res == 2586


@pytest.mark.benchmark(group="day25")
def test_part2(benchmark):
    res = benchmark(part2)
    assert res is True
