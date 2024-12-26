from collections import defaultdict
from itertools import pairwise
from functools import cache


def processing() -> list[int]:
    return list(map(int, open("src/day22/input.txt").read().splitlines()))


@cache
def calculate(secret: int, n: int = 1) -> int:
    if n == 1:
        secret = (secret << 6 ^ secret) & 0xFFFFFF
        secret = (secret >> 5 ^ secret) & 0xFFFFFF
        return (secret << 11 ^ secret) & 0xFFFFFF
    for _ in range(n):
        secret = (secret << 6 ^ secret) & 0xFFFFFF
        secret = (secret >> 5 ^ secret) & 0xFFFFFF
        secret = (secret << 11 ^ secret) & 0xFFFFFF
    return secret


def part1() -> int:
    return sum(calculate(secret, n=2000) for secret in processing())


def part2() -> int:
    combinations: dict[tuple, int] = defaultdict(int)
    for secret in processing():
        sequence: list[int] = [secret] + [secret := calculate(secret) for _ in range(2000)]
        differences: list[int] = [second % 10 - first % 10 for first, second in pairwise(sequence)]
        seen: set[tuple] = set()
        for i in range(len(sequence)-4):
            pattern: tuple = tuple(differences[i:i+4])
            if pattern not in seen:
                combinations[pattern] += sequence[i+4] % 10
                seen.add(pattern)
    return max(combinations.values())


if __name__ == "__main__":
    print(part1())
    print(part2())
