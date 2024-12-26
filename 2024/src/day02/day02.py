def processing() -> list[list[int]]:
    return [list(map(int, level.split())) for level in open("src/day02/input.txt")]


def check_safe(level: list, tolerance: int = 0) -> bool:
    for i in range(len(level)-1):
        if not 1 <= level[i] - level[i+1] <= 3:
            return tolerance and any(check_safe(level[j-1:j] + level[j+1:]) for j in (i, i+1))
    return True


def part1() -> int:
    return sum([check_safe(level) or check_safe(level[::-1]) for level in processing()])


def part2() -> int:
    return sum([check_safe(level, 1) or check_safe(level[::-1], 1) for level in processing()])


if __name__ == "__main__":
    print(part1())
    print(part2())
