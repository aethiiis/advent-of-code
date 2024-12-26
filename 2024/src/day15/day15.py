def processing():
    top, bottom = open("src/day15/input.txt").read().split("\n\n")
    return top, bottom.replace("\n", "")


def part1() -> int:
    top, commands = processing()
    grid = [list(line) for line in top.splitlines()]
    rows, cols = len(grid), len(grid[0])
    r, c = [(r, c) for r in range(rows) for c in range(cols) if grid[r][c] == "@"][0]
    for command in commands:
        dr, dc = {"^": -1, "v": 1}.get(command, 0), {"<": -1, ">": 1}.get(command, 0)
        moved, movable = [(r, c)], True
        cr, cc = r, c
        while True:
            cr += dr
            cc += dc
            value = grid[cr][cc]
            if value == "#":
                movable = False
                break
            if value == "O":
                moved.append((cr, cc))
            if value == ".":
                break
        if not movable:
            continue
        grid[r][c], grid[r + dr][c + dc] = ".", "@"
        for br, bc in moved[1:]:
            grid[br + dr][bc + dc] = "O"
        r += dr
        c += dc
    return sum(100 * r + c for r in range(rows) for c in range(cols) if grid[r][c] == "O")


def part2() -> int:
    top, commands = processing()
    expansion = {"#": "##", "O": "[]", ".": "..", "@": "@."}
    grid = [list("".join(expansion[char] for char in line)) for line in top.splitlines()]
    rows, cols = len(grid), len(grid[0])
    r, c = [(r, c) for r in range(rows) for c in range(cols) if grid[r][c] == "@"][0]
    for command in commands:
        dr, dc = {"^": -1, "v": 1}.get(command, 0), {"<": -1, ">": 1}.get(command, 0)
        moved, movable = [(r, c)], True
        for cr, cc in moved:
            nr, nc = cr + dr, cc + dc
            if (nr, nc) in moved:
                continue
            char = grid[nr][nc]
            if char == "#":
                movable = False
                break
            if char == "[":
                moved.append((nr, nc))
                moved.append((nr, nc + 1))
            if char == "]":
                moved.append((nr, nc))
                moved.append((nr, nc - 1))
        if not movable:
            continue
        copy = [list(row) for row in grid]
        grid[r][c], grid[r + dr][c + dc] = ".", "@"
        for br, bc in moved[1:]:
            grid[br][bc] = "."
        for br, bc in moved[1:]:
            grid[br + dr][bc + dc] = copy[br][bc]
        r += dr
        c += dc
    return sum(100 * r + c for r in range(rows) for c in range(cols) if grid[r][c] == "[")


if __name__ == "__main__":
    print(part1())
    print(part2())
