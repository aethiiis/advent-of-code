from bisect import bisect


def processing():
    bits = [tuple(map(int, bit.split(","))) for bit in open("src/day18/input.txt").read().splitlines()]
    grid = [["." for _ in range(71)] for _ in range(71)]
    return grid, bits


def run(bits, i, start, end, rows=71, cols=71):
    seen = {*bits[:i]}
    frontier = [(0, start)]
    for cost, (r, c) in frontier:
        if (r, c) == end:
            return cost
        for nr, nc in [(r, c+1), (r+1, c), (r-1, c), (r, c-1)]:
            if (nr, nc) not in seen and 0 <= nr < rows and 0 <= nc < cols:
                frontier.append((cost+1, (nr, nc)))
                seen.add((nr, nc))
    return 1e9


def part1() -> str:
    return run(processing()[1], 1024, (0, 0), (70, 70))


def part2() -> str:
    grid, bits = processing()
    return ",".join(map(str, bits[bisect(range(len(bits)), int(1e9-1), key=lambda i: run(bits, i, (0, 0), (70, 70)))-1]))


if __name__ == "__main__":
    print(part1())
    print(part2())
