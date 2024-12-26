import pytest
from day02 import part1, part2


@pytest.mark.benchmark(group="day02")
def test_part1(benchmark):
    res = benchmark(part1)
    assert res == 663


@pytest.mark.benchmark(group="day02")
def test_part2(benchmark):
    res = benchmark(part2)
    assert res == 692
