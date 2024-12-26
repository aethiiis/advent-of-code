import pytest
from day18 import part1, part2


@pytest.mark.benchmark(group="day18")
def test_part1(benchmark):
    res = benchmark(part1)
    assert res == 278


@pytest.mark.benchmark(group="day18")
def test_part2(benchmark):
    res = benchmark(part2)
    assert res == "43,12"
