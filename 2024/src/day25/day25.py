import os


def processing():
    blocks = list(map(lambda x: x.splitlines(), open("src/day25/input.txt").read().split("\n\n")))
    rows = len(blocks[0])
    cols = len(blocks[0][0])
    keys = [block for block in blocks if block[0] != "#####"]
    locks = [block for block in blocks if block[0] == "#####"]
    keys_height = {}
    locks_height = {}
    for i, key in enumerate(keys):
        combination = [0, 0, 0, 0, 0]
        for c in range(cols):
            column: int = rows - 1
            for r in range(rows):
                if key[r][c] == "#":
                    break
                else:
                    column -= 1
            combination[c] = column
        keys_height[i] = combination
    for i, lock in enumerate(locks):
        combination = [0, 0, 0, 0, 0]
        for c in range(cols):
            column: int = rows - 1
            for r in range(rows - 1, -1, -1):
                if lock[r][c] == "#":
                    break
                else:
                    column -= 1
            combination[c] = column
        locks_height[i] = combination
    return keys_height, locks_height


def part1():
    keys_height, locks_height = processing()
    count = 0
    for i, key in keys_height.items():
        for j, lock in locks_height.items():
            for k in range(5):
                if key[k] + lock[k] > 5:
                    break
            else:
                count += 1
    return count


def part2():
    for i in range(1, 26):
        number = str(i).rjust(2, "0")
        path = "src/day" + number
        if (not os.path.exists(path) or
                not os.path.isfile(os.path.join(path, "input.txt")) or
                not os.path.isfile(os.path.join(path, "test.txt")) or
                not os.path.isfile(os.path.join(path, "day" + number + ".py")) or
                not os.path.isfile(os.path.join(path, "test_day" + number + ".py"))):
            return False
    return True


if __name__ == '__main__':
    print(part1())
    print(part2())
