import pytest
from day22 import part1, part2


@pytest.mark.benchmark(group="day22")
def test_part1(benchmark):
    res = benchmark(part1)
    assert res == 14726157693


@pytest.mark.benchmark(group="day22")
def test_part2(benchmark):
    res = benchmark(part2)
    assert res == 1614
