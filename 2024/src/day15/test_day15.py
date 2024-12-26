import pytest
from day15 import part1, part2


@pytest.mark.benchmark(group="day15")
def test_part1(benchmark):
    res = benchmark(part1)
    assert res == 1478649


@pytest.mark.benchmark(group="day15")
def test_part2(benchmark):
    res = benchmark(part2)
    assert res == 1495455
