import pytest
from day08 import part1, part2


@pytest.mark.benchmark(group="day08")
def test_part1(benchmark):
    res = benchmark(part1)
    assert res == 361


@pytest.mark.benchmark(group="day08")
def test_part2(benchmark):
    res = benchmark(part2)
    assert res == 1249
