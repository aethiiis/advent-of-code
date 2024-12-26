from queue import Queue
from functools import cache
from itertools import product


def processing() -> tuple[list[str], list[list[str | None]], list[list[str | None]]]:
    return (open("src/day21/input.txt").read().splitlines(), [[None, "^", "A"], ["<", "v", ">"]],
            [["7", "8", "9"], ["4", "5", "6"], ["1", "2", "3"], [None, "0", "A"]])


def postprocessing() -> tuple[
        list[str], dict[tuple[str, str], list[str]], dict[tuple[str, str], list[str]], dict[tuple[str, str], int]]:
    codes, dir_kp, num_kp = processing()
    dir_sequences = get_sequences(dir_kp)
    dir_sequences_length: dict[tuple[str, str], int] = {step: len(path[0]) for step, path in dir_sequences.items()}
    num_sequences = get_sequences(num_kp)
    return codes, num_sequences, dir_sequences, dir_sequences_length


def get_positions(keypad: list[list[str | None]]) -> dict[str, tuple[int, int]]:
    return {keypad[r][c]: (r, c) for r in range(len(keypad)) for c in range(len(keypad[0])) if keypad[r][c] is not None}


def get_sequences(keypad: list[list[str | None]]) -> dict[tuple[str, str], list[str]]:
    positions = get_positions(keypad)
    sequences: dict[tuple[str, str], list[str]] = {}
    for start in positions:
        for end in positions:
            if start == end:
                sequences[(start, end)] = ["A"]
                continue
            possibilities = bfs(keypad, positions, start, end)
            if possibilities is None:
                print(start, end, )
            sequences[(start, end)] = possibilities
    return sequences


def bfs(keypad: list[list[str | None]], positions: dict[str, tuple[int, int]], start: str, end: str) -> list[str]:
    frontier = Queue()
    frontier.put((positions[start], ""))
    cost: float = float('inf')
    possibilities: list[str] = []
    while not frontier.empty():
        (r, c), path = frontier.get()
        for dr, dc, direction in [(0, 1, ">"), (1, 0, "v"), (0, -1, "<"), (-1, 0, "^")]:
            nr, nc = r + dr, c + dc
            if 0 <= nr < len(keypad) and 0 <= nc < len(keypad[0]) and keypad[nr][nc] is not None:
                if keypad[nr][nc] == end:
                    new_cost = len(path) + 1
                    if cost < new_cost:
                        return possibilities
                    cost = new_cost
                    possibilities.append(path + direction + "A")
                else:
                    frontier.put(((nr, nc), path + direction))
    return possibilities


def run(code: str, sequences: dict[tuple[str, str], list[str]]) -> list[str]:
    return ["".join(possibility) for possibility in
            product(*[sequences[(start, end)] for start, end in zip("A" + code, code)])]


def part1():
    codes, num_sequences, dir_sequences, dir_sequences_length = postprocessing()

    @cache
    def solve_part1(sequence: str, depth=2) -> int:
        if depth == 1:
            return sum(dir_sequences_length[(start, end)] for start, end in zip("A" + sequence, sequence))
        len_ = 0
        for start, end in zip("A" + sequence, sequence):
            len_ += min(solve_part1(subsequence, depth - 1) for subsequence in dir_sequences[(start, end)])
        return len_

    count: int = 0
    for code in codes:
        possibilities = run(code, num_sequences)
        count += min(map(solve_part1, possibilities)) * int(code[:-1])
    return count


def part2():
    codes, num_sequences, dir_sequences, dir_sequences_length = postprocessing()

    @cache
    def solve_part2(sequence: str, depth=25) -> int:
        if depth == 1:
            return sum(dir_sequences_length[(start, end)] for start, end in zip("A" + sequence, sequence))
        len_ = 0
        for start, end in zip("A" + sequence, sequence):
            len_ += min(solve_part2(subsequence, depth - 1) for subsequence in dir_sequences[(start, end)])
        return len_

    count: int = 0
    for code in codes:
        possibilities = run(code, num_sequences)
        count += min(map(solve_part2, possibilities)) * int(code[:-1])
    return count


if __name__ == "__main__":
    print(part1())
    print(part2())
