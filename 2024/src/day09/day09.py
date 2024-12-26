def processing():
    return open("src/day09/input.txt").read()


def part1() -> int:
    disk: list[int] = [i // 2 if i % 2 == 0 else -1 for i, char in enumerate(processing()) for _ in range(int(char))]
    blanks = [i for i, x in enumerate(disk) if x == -1]
    for i in blanks:
        while disk[-1] == -1:
            disk.pop()
        if len(disk) <= i:
            break
        disk[i] = disk.pop()
    return sum(i * x for i, x in enumerate(disk))


def part2() -> int:
    files: dict[int, tuple[int, int]] = {}
    blanks: list[tuple[int, int]] = []
    fid, pos = 0, 0
    for i, char in enumerate(processing()):
        x = int(char)
        if i % 2 == 0:
            files[fid] = (pos, x)
            fid += 1
        else:
            if x != 0:
                blanks.append((pos, x))
        pos += x
    while fid > 0:
        fid -= 1
        pos, size = files[fid]
        for i, (start, length) in enumerate(blanks):
            if start >= pos:
                blanks = blanks[:i]
                break
            if size <= length:
                files[fid] = (start, size)
                if size == length:
                    blanks.pop(i)
                else:
                    blanks[i] = (start + size, length - size)
                break
    return sum(fid * size * (2 * pos + size - 1) // 2 for fid, (pos, size) in files.items())


if __name__ == "__main__":
    print(part1())
    print(part2())
