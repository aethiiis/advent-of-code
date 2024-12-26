from functools import cache


def processing() -> dict[complex, str]:
    return {i + j * 1j: c for i, r in enumerate(open("src/day06/input.txt").readlines()) for j, c in
            enumerate(r.strip())}


def print_grid(grid: dict[complex, str]) -> None:
    for i in range(10):
        for j in range(10):
            print(grid[(i + j * 1j)], end="")
        print()


def out(grid: dict[complex, str], start: complex) -> set[complex]:
    p, d, visited = start, -1, set()
    while p in grid and (p, d) not in visited:
        visited.add((p, d))
        if grid.get(p + d) == "#":
            d *= -1j
        else:
            p += d
    return {p for p, _ in visited}


def part1() -> int:
    grid: dict[complex, str] = processing()
    return len(out(grid, [p for p in grid if grid[p] == '^'][0]))


@cache
def part2() -> int:
    start_grid: dict[complex, str] = processing()
    start: complex = [p for p in start_grid if start_grid[p] == '^'][0]
    count = 0
    for o in out(start_grid, start):
        if o == start:
            continue
        grid = start_grid | {o: "#"}
        p, d, visited = start, -1, set()
        while p in grid and (p, d) not in visited:
            visited.add((p, d))
            if grid.get(p + d) == "#":
                d *= -1j
            else:
                p += d
        count += (p, d) in visited
    return count


if __name__ == "__main__":
    print(part1())
    print(part2())
