def processing() -> list:
    return [list(map(lambda slist: tuple(map(int, slist)), map(lambda x: x.split("=")[1].split(","), line.split(" "))))
            for line in open("src/day14/input.txt").read().splitlines()]


def detect_repeats(robots, w=101, h=103):
    first, delta = [0, 0], [0, 0]
    i = 0
    while True:
        i += 1
        positions = [((x + vx * i) % w, (y + vy * i) % h) for (x, y), (vx, vy) in robots]
        for j in [0, 1]:
            if detect_lines(positions, j):
                if first[j] == 0:
                    first[j] = i
                elif delta[j] == 0:
                    delta[j] = i - first[j]
        if all(d != 0 for d in delta):
            break

    return first, delta


def detect_lines(positions, mode):
    if mode == 0:
        positions.sort(key=lambda p: (p[0], p[1]))
    else:
        positions.sort(key=lambda p: (p[1], p[0]))
    count = 0
    old_p = positions[0]
    for i in range(1, len(positions)):
        if positions[i][mode] != old_p[mode]:
            count = 0
        elif positions[i][(mode + 1) % 2] - old_p[(mode + 1) % 2] <= 2:
            count += 1
            if count >= 10:
                return True
        old_p = positions[i]
    return False


def extended_euclidean(a: int, b: int):
    r1, u1, v1, r2, u2, v2 = a, 1, 0, b, 0, 1
    while r2 != 0:
        q = r1 // r2
        r1, u1, v1, r2, u2, v2 = r2, u2, v2, r1 - q * r2, u1 - q * u2, v1 - q * v2
    return r1, u1, v1


def chinese_remainder(first, delta):
    _, u, v = extended_euclidean(delta[1], delta[0])
    return (first[1] * delta[0] * v + first[0] * delta[1] * u) % (delta[1] * delta[0])


def run(robots, t, w=101, h=103) -> int:
    a = b = c = d = 0
    for (x, y), (vx, vy) in robots:
        x, y = (x + vx * t) % w, (y + vy * t) % h
        a += x > w // 2 and y > h // 2
        b += x > w // 2 and y < h // 2
        c += x < w // 2 and y > h // 2
        d += x < w // 2 and y < h // 2
    return a * b * c * d


def part1() -> int:
    return run(processing(), 100)


def part2() -> int:
    return chinese_remainder(*detect_repeats(processing()))


if __name__ == "__main__":
    print(part1())
    print(part2())
