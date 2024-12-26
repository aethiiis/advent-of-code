from functools import cmp_to_key


def processing() -> list[int]:
    rules, pages = open('src/day05/input.txt').read().split('\n\n')
    cmp = cmp_to_key(lambda x, y: -(x + '|' + y in rules))
    count: list[int] = [0, 0]
    for p in pages.split():
        p = p.split(',')
        s = sorted(p, key=cmp)
        count[p != s] += int(s[len(s) // 2])
    return count


def part1() -> int:
    return processing()[0]


def part2() -> int:
    return processing()[1]


if __name__ == "__main__":
    print(part1())
    print(part2())
