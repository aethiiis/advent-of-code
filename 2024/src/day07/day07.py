def processing() -> list[list[int | list[int]]]:
    return [[int(equation[0]), list(map(int, equation[1].split()))]
            for equation in (line.split(": ")
                             for line in open("src/day07/input.txt").read().splitlines())]


def obtainable(result, numbers, concat) -> bool:
    if len(numbers) == 1:
        return result == numbers[0]
    if result % numbers[-1] == 0 and obtainable(result // numbers[-1], numbers[:-1], concat):
        return True
    if result > numbers[-1] and obtainable(result - numbers[-1], numbers[:-1], concat):
        return True
    if concat:
        r, n = str(result), str(numbers[-1])
        if len(r) > len(n) and r.endswith(n) and obtainable(int(r[:-len(n)]), numbers[:-1], concat):
            return True


def part1() -> int:
    return sum(equation[0] for equation in processing() if obtainable(equation[0], equation[1], False))


def part2() -> int:
    return sum(equation[0] for equation in processing() if obtainable(equation[0], equation[1], True))


if __name__ == "__main__":
    print(part1())
    print(part2())
