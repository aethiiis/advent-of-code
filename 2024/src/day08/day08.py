from itertools import product


def processing():
    grid = [list(line) for line in open("src/day08/input.txt").read().splitlines()]
    return grid


def calculate_antinodes(node1, node2):
    return (node1[0] - (node2[0] - node1[0]), node1[1] - (node2[1] - node1[1])), (
        node2[0] + (node2[0] - node1[0]), node2[1] + (node2[1] - node1[1]))


def calculate_resonance(node1, node2, grid):
    return [x for i in range(max(len(grid)//(node2[0]-node1[0]), len(grid[0])//(node2[1]-node1[1])))
            for x in [(node1[0]-i*(node2[0]-node1[0]), node1[1]-i*(node2[1]-node1[1])),
                      (node2[0]+i*(node2[0]-node1[0]), node2[1]+i*(node2[1]-node1[1]))]]


def find_frequencies(grid):
    return {node: [(r, c) for r, row in enumerate(grid) for c, n in enumerate(row) if n == node]
            for node in {n for row in grid for n in row if n != "."}}


def is_in(node, grid):
    return 0 <= node[0] <= len(grid) - 1 and 0 <= node[1] <= len(grid) - 1


def part1():
    grid = processing()
    return len({antinode for antennas in find_frequencies(grid).values() for node1, node2 in product(antennas, repeat=2)
                if node1 != node2 for antinode in calculate_antinodes(node1, node2) if is_in(antinode, grid)})


def part2():
    grid = processing()
    return len({antinode for antennas in find_frequencies(grid).values() for node1, node2 in product(antennas, repeat=2)
                if node1 != node2 for antinode in calculate_resonance(node1, node2, grid) if is_in(antinode, grid)})


if __name__ == "__main__":
    print(part1())
    print(part2())
