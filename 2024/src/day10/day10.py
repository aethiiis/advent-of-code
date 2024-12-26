from queue import Queue


def processing():
    lines = [[int(x) for x in line] for line in open("src/day10/input.txt").read().split()]
    rows = len(lines)
    cols = len(lines[0])
    return {(r, c): lines[r][c] for r in range(rows) for c in range(cols)}, rows, cols


def search(topo, rows, cols, zeroes, memory=True):
    trailheads = dict()
    for zero in zeroes:
        frontier = Queue()
        frontier.put(zero)
        if memory:
            nines = set()
        while not frontier.empty():
            (r, c) = frontier.get()
            value = topo[(r, c)]
            if value == 9:
                if memory and (r, c) not in nines:
                    nines.add((r, c))
                    if zero not in trailheads:
                        trailheads[zero] = 1
                    else:
                        trailheads[zero] += 1
                elif memory:
                    continue
                else:
                    if zero not in trailheads:
                        trailheads[zero] = 1
                    else:
                        trailheads[zero] += 1
            else:
                valid = neighbors(r, c, rows, cols, topo, value)
                for x in valid:
                    frontier.put(x)
    return trailheads


def neighbors(r, c, rows, cols, topo, value):
    return [(r+dr, c+dc) for (dr, dc) in [(-1, 0), (1, 0), (0, -1), (0, 1)]
            if 0 <= r+dr <= rows-1 and 0 <= c+dc <= cols-1 and topo[(r+dr, c+dc)] == value+1]


def part1() -> int:
    topo, rows, cols = processing()
    return sum(x for x in search(topo, rows, cols, [(r, c) for (r, c), x in topo.items() if x == 0]).values())


def part2() -> int:
    topo, rows, cols = processing()
    return sum(x for x in search(topo, rows, cols, [(r, c) for (r, c), x in topo.items() if x == 0], False).values())


if __name__ == "__main__":
    print(part1())
    print(part2())
