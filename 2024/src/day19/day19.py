from functools import cache


def processing() -> tuple[frozenset, list[str]]:
    patterns, towels = open("src/day19/input.txt").read().split('\n\n')
    return frozenset([pattern for pattern in patterns.split(", ")]), towels.splitlines()


@cache
def possibilities(patterns, towel: str):
    if not towel:
        return 1
    return sum(possibilities(patterns, towel[len(p):]) for p in patterns if towel.startswith(p))


def part1():
    patterns, towels = processing()
    return sum(1 for towel in towels if possibilities(patterns, towel) > 1)


def part2():
    patterns, towels = processing()
    return sum(possibilities(patterns, towel) for towel in towels)


if __name__ == "__main__":
    print(part1())
    print(part2())
