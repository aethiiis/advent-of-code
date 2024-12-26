import pytest
from day20 import part1, part2


@pytest.mark.benchmark(group="day20")
def test_part1(benchmark):
    res = benchmark(part1)
    assert res == 1402


@pytest.mark.benchmark(group="day20")
def test_part2(benchmark):
    res = benchmark(part2)
    assert res == 1020244
