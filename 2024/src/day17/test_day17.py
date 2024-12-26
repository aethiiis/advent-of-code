import pytest
from day17 import part1, part2


@pytest.mark.benchmark(group="day17")
def test_part1(benchmark):
    res = benchmark(part1)
    assert res == "1,2,3,1,3,2,5,3,1"


@pytest.mark.benchmark(group="day17")
def test_part2(benchmark):
    res = benchmark(part2)
    assert res == 105706277661082
