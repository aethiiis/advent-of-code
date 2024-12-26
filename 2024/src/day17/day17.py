def processing():
    registers, program = open("src/day17/input.txt").read().split("\n\n")
    registers = {register[9:10]: int(register[12:]) for register in registers.split("\n")}
    program = list(map(int, program[9:].split(",")))
    return registers, program


class Computer:
    def __init__(self, registers: dict, program: list) -> None:
        self.registers = registers
        self.program = program
        self.pointer = 0
        self.output = []

    def combo(self, operand: int) -> int:
        if 0 <= operand <= 3:
            return operand
        elif operand == 4:
            return self.registers["A"]
        elif operand == 5:
            return self.registers["B"]
        elif operand == 6:
            return self.registers["C"]

    def _adv(self, operand) -> None:
        self.registers["A"] = self.registers["A"] >> self.combo(operand)

    def _bxl(self, operand) -> None:
        self.registers["B"] = self.registers["B"] ^ operand

    def _bst(self, operand) -> None:
        self.registers["B"] = self.combo(operand) % 8

    def _jnz(self, operand) -> None:
        if self.registers["A"] != 0:
            self.pointer = operand - 2

    def _bxc(self) -> None:
        self.registers["B"] = self.registers["B"] ^ self.registers["C"]

    def _out(self, operand) -> None:
        self.output.append(self.combo(operand) % 8)

    def _bdv(self, operand) -> None:
        self.registers["B"] = self.registers["A"] >> self.combo(operand)

    def _cdv(self, operand) -> None:
        self.registers["C"] = self.registers["A"] >> self.combo(operand)

    def run(self) -> None:
        while self.pointer < len(self.program):
            code, op = self.program[self.pointer:self.pointer + 2]
            match code, op:
                case 0, op:
                    self._adv(op)
                case 1, op:
                    self._bxl(op)
                case 2, op:
                    self._bst(op)
                case 3, op:
                    self._jnz(op)
                case 4, _:
                    self._bxc()
                case 5, op:
                    self._out(op)
                case 6, op:
                    self._bdv(op)
                case 7, op:
                    self._cdv(op)
            self.pointer += 2


def find(program, value):
    if len(program) == 0:
        return value
    for i in range(8):
        a = value << 3 | i
        b = a % 8
        b = b ^ 5
        c = a >> b
        b = b ^ 6
        b = b ^ c
        if b % 8 == program[-1]:
            res = find(program[:-1], a)
            if res is None:
                continue
            return res


def part1() -> str:
    c = Computer(*processing())
    c.run()
    return ",".join(list(map(str, c.output)))


def part2() -> int:
    return find(processing()[1], 0)


if __name__ == "__main__":
    print(part1())
    print(part2())
