import pytest
from day24 import part1, part2


@pytest.mark.benchmark(group="day24")
def test_part1(benchmark):
    res = benchmark(part1)
    assert res == 59619940979346


@pytest.mark.benchmark(group="day24")
def test_part2(benchmark):
    res = benchmark(part2)
    assert res == "bpt,fkp,krj,mfm,ngr,z06,z11,z31"
