import pytest
from day14 import part1, part2


@pytest.mark.benchmark(group="day14")
def test_part1(benchmark):
    res = benchmark(part1)
    assert res == 218965032


@pytest.mark.benchmark(group="day14")
def test_part2(benchmark):
    res = benchmark(part2)
    assert res == 7037
