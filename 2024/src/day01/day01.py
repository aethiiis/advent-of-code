from collections import Counter


def processing() -> tuple[list[int], list[int]]:
    lines: list[list[int]] = [list(map(int, line.split())) for line in open("src/day01/input.txt")]
    return [lines[i][1] for i in range(len(lines))], [lines[i][0] for i in range(len(lines))]


def part1() -> int:
    left, right = map(sorted, processing())
    return sum([abs(left[i] - right[i]) for i in range(len(left))])


def part2() -> int:
    # noinspection PyTypeChecker
    left, right = map(Counter, processing())
    return sum([0 if key not in right else key * value * right[key] for key, value in left.items()])


if __name__ == "__main__":
    print(part1())
    print(part2())
