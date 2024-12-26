def processing():
    grid = [list(line) for line in open("src/day20/input.txt").read().splitlines()]
    rows = len(grid)
    cols = len(grid[0])
    limits = [(r, c) for r in range(rows) for c in range(cols) if grid[r][c] == "S" or grid[r][c] == "E"]
    return grid, rows, cols, limits[0], limits[1]


def course(grid, rows, cols, start, end):
    path = [start]
    distances = [[-1 for _ in range(cols)] for _ in range(rows)]
    distances[start[0]][start[1]] = 0
    position = start
    while position != end:
        for dr, dc in [(1, 0), (-1, 0), (0, 1), (0, -1)]:
            r, c = position[0] + dr, position[1] + dc
            if 0 <= r < rows and 0 <= c < cols and grid[r][c] != "#" and distances[r][c] == -1:
                distances[r][c] = distances[position[0]][position[1]] + 1
                path.append((r, c))
                position = (r, c)
    return distances, path


def part1():
    grid, rows, cols, start, end = processing()
    distances, path = course(grid, rows, cols, start, end)
    return sum(1 if 0 < nr < rows - 1 and 0 < nc < cols - 1 and grid[nr][nc] != "#" and
                    abs(distances[r][c] - distances[nr][nc]) >= 102 else 0
               for (r, c) in path for (nr, nc) in [(r + 2, c), (r + 1, c + 1), (r, c + 2), (r - 1, c + 1)])


def part2():
    grid, rows, cols, start, end = processing()
    distances, path = course(grid, rows, cols, start, end)
    return sum(1 if 0 < nr < rows-1 and 0 < nc < cols-1 and grid[nr][nc] != "#" and
                    distances[r][c] - distances[nr][nc] >= radius + 100 else 0
               for (r, c) in path for radius in range(2, 21) for dr in range(radius+1)
               for (nr, nc) in {(r+dr, c+radius-dr), (r+dr, c-radius+dr), (r-dr, c+radius-dr), (r-dr, c-radius+dr)})


if __name__ == "__main__":
    print(part1())
    print(part2())
