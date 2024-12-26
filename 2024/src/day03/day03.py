from re import findall


def processing() -> list:
    return [0 if match == "don't()" else
            1 if match == "do()" else
            [int(num) for num in match[4:-1].split(",")] for match in
            findall(r"do\(\)|don't\(\)|mul\(\d+,\d+\)",
                    "".join(open("src/day03/input.txt", "r").readlines()))]


def part1() -> int:
    return sum([couple[0] * couple[1] for couple in processing() if couple != 0 and couple != 1])


def part2() -> int:
    nums: list = processing()
    enabled: bool = True
    count: int = 0
    for couple in nums:
        if couple == 0:
            enabled = False
        elif couple == 1:
            enabled = True
        elif enabled:
            count += couple[0] * couple[1]
    return count


if __name__ == "__main__":
    print(part1())
    print(part2())
