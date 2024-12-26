import pytest
from day05 import part1, part2


@pytest.mark.benchmark(group="day05")
def test_part1(benchmark):
    res = benchmark(part1)
    assert res == 6242


@pytest.mark.benchmark(group="day05")
def test_part2(benchmark):
    res = benchmark(part2)
    assert res == 5169
