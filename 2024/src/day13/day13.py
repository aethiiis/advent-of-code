def processing() -> list[dict[str, tuple[int, int]]]:
    return [{
        "A": (int(line[2][2:-1]), int(line[3][2:])),
        "B": (int(line[6][2:-1]), int(line[7][2:])),
        "Target": (int(line[9][2:-1]), int(line[10][2:]))
    } for line in [line.split() for line in open("src/day13/input.txt").read().split("\n\n")]]


def solve(machines) -> int:
    count = 0
    for machine in machines:
        ax, ay = machine["A"]
        bx, by = machine["B"]
        tx, ty = machine["Target"]
        a = (tx * by - ty * bx) // (ax * by - ay * bx)
        b = (tx * ay - ty * ax) // (ay * bx - ax * by)
        if ax * a + bx * b == tx and ay * a + by * b == ty:
            count += 3 * a + b
    return count


def part1() -> int:
    return solve(processing())


def part2() -> int:
    machines = processing()
    for machine in machines:
        machine["Target"] = (machine["Target"][0] + 10000000000000, machine["Target"][1] + 10000000000000)
    return solve(machines)


if __name__ == "__main__":
    print(part1())
    print(part2())
