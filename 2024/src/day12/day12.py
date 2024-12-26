from numpy import array, unique
from scipy.ndimage import label
from scipy.signal import convolve2d


def processing() -> array:
    return array([list(line) for line in open('src/day12/input.txt').read().splitlines()])


def arrsum(arr) -> int:
    return abs(arr).sum()


def part1() -> int:
    farm = processing()
    count = 0
    for garden, number in [label(farm == plant) for plant in unique(farm)]:
        for i in range(number):
            plot = (garden == i + 1)
            count += arrsum(plot) * (arrsum(convolve2d(plot, [[1, -1]])) + arrsum(convolve2d(plot, [[1], [-1]])))
    return count


def part2() -> int:
    farm = processing()
    count = 0
    for garden, number in [label(farm == plant) for plant in unique(farm)]:
        for i in range(number):
            plot = (garden == i + 1)
            count += arrsum(plot) * arrsum(convolve2d(plot, [[-1, 1], [1, -1]]))
    return count


if __name__ == "__main__":
    print(part1())
    print(part2())
