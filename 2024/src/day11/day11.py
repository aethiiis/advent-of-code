from functools import cache


def processing() -> list[int]:
    stones = list(map(int, open("src/day11/input.txt").read().split(" ")))
    return stones


@cache
def handle_stone(stone, times) -> int:
    if not times:
        return 1
    elif not stone:
        return handle_stone(1, times - 1)
    elif (s := str(stone)) and not (le := len(s)) % 2:
        return handle_stone(int(s[:le // 2]), times - 1) + handle_stone(int(s[le // 2:]), times - 1)
    else:
        return handle_stone(stone * 2024, times - 1)


def part1() -> int:
    return sum(handle_stone(stone, 25) for stone in processing())


def part2() -> int:
    return sum(handle_stone(stone, 75) for stone in processing())


if __name__ == "__main__":
    print(part1())
    print(part2())
