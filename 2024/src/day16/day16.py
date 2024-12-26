from queue import PriorityQueue


def processing():
    i = open("src/day16/input.txt").read().splitlines()
    grid = {(r, c): v for r in range(len(i)) for c, v in enumerate(i[r])}
    return grid, len(i), len(i[0])


def dijkstra(grid, starts, rows, cols):
    frontier = PriorityQueue()
    best = dict()
    for start in starts:
        node, direction = start
        frontier.put((0, node, direction))
    while not frontier.empty():
        cost, node, direction = frontier.get()
        if (node, direction) in best and cost >= best[(node, direction)]:
            continue
        best[(node, direction)] = cost
        for dr, dc in [direction, (direction[1], -direction[0]), (-direction[1], direction[0])]:
            r, c = node[0] + dr, node[1] + dc
            if 0 <= r < rows and 0 <= c < cols and grid[(r, c)] in ['.', 'S', 'E'] and (dr, dc) == direction:
                frontier.put((cost + 1, (r, c), (dr, dc)))
            if (dr, dc) != direction:
                frontier.put((cost + 1000, node, (dr, dc)))
    return best


def part1() -> int:
    grid, rows, cols = processing()
    start, end = (rows - 2, 1), (1, cols - 2)
    paths = dijkstra(grid, [[start, (0, 1)]], rows, cols)
    return min(paths[(end, (dr, dc))] for dr, dc in [(0, 1), (1, 0), (0, -1), (-1, 0)])


def part2() -> int:
    grid, rows, cols = processing()
    start, end = (rows - 2, 1), (1, cols - 2)
    from_start, from_end = [dijkstra(grid, begin, rows, cols) for begin in
                            ([[start, (0, 1)]], [[end, (dr, dc)] for dr, dc in [(0, 1), (1, 0), (0, -1), (-1, 0)]])]
    return len({(r, c) for r in range(rows) for c in range(cols) for dr, dc in [(0, 1), (1, 0), (0, -1), (-1, 0)] if
                from_start.get(((r, c), (dr, dc)), 0) + from_end.get(((r, c), (-dr, -dc)), 0) ==
                min(from_start[(end, (dr, dc))] for dr, dc in [(0, 1), (1, 0), (0, -1), (-1, 0)])})


if __name__ == "__main__":
    print(part1())
    print(part2())
