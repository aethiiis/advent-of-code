import pytest
from day16 import part1, part2


@pytest.mark.benchmark(group="day16")
def test_part1(benchmark):
    res = benchmark(part1)
    assert res == 111480


@pytest.mark.benchmark(group="day16")
def test_part2(benchmark):
    res = benchmark(part2)
    assert res == 529
