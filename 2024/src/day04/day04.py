def processing() -> list[str]:
    return [line.strip() for line in open("src/day04/input.txt").readlines()]


def part1() -> int:
    grid = processing()
    rows = len(grid)
    cols = len(grid[0])
    return sum([all(
        0 <= row + dr * i < rows and 0 <= col + dc * i < cols and grid[row + dr * i][col + dc * i] == "XMAS"[i] for i in
        range(4)) for row in range(rows) for col in range(cols) for dr, dc in
        [(1, 0), (0, 1), (-1, 0), (0, -1), (1, 1), (-1, -1), (1, -1), (-1, 1)]])


def part2() -> int:
    grid = processing()
    return sum([grid[row][col] == "A" and
                ((grid[row - 1][col - 1] == "M" and grid[row + 1][col - 1] == "S" and grid[row - 1][col + 1] == "M" and
                  grid[row + 1][col + 1] == "S") or
                 (grid[row - 1][col - 1] == "S" and grid[row + 1][col - 1] == "M" and grid[row - 1][col + 1] == "S" and
                  grid[row + 1][col + 1] == "M") or
                 (grid[row - 1][col - 1] == "M" and grid[row + 1][col - 1] == "M" and grid[row - 1][col + 1] == "S" and
                  grid[row + 1][col + 1] == "S") or
                 (grid[row - 1][col - 1] == "S" and grid[row + 1][col - 1] == "S" and grid[row - 1][col + 1] == "M" and
                  grid[row + 1][col + 1] == "M"))
                for row in range(1, len(grid) - 1) for col in range(1, len(grid[0]) - 1)])


if __name__ == "__main__":
    print(part1())
    print(part2())
