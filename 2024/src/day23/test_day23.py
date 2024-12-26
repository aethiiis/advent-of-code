import pytest
from day23 import part1, part2


@pytest.mark.benchmark(group="day23")
def test_part1(benchmark):
    res = benchmark(part1)
    assert res == 1366


@pytest.mark.benchmark(group="day23")
def test_part2(benchmark):
    res = benchmark(part2)
    assert res == "bs,cf,cn,gb,gk,jf,mp,qk,qo,st,ti,uc,xw"
