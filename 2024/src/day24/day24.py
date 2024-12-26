from graphviz import Digraph


def processing():
    w, g = open("src/day24/input.txt").read().split("\n\n")
    return ({wire: int(value) for wire, value in map(lambda x: x.split(": "), w.splitlines())},
            {gate[4]: (gate[0], gate[1], gate[2]) for gate in map(lambda x: x.split(" "), g.splitlines())})


def checks(gates):
    wrong = set()
    for wire, (left, op, right) in gates.items():
        if wire[0] == "z" and op != "XOR" and wire != "z45":
            wrong.add(wire)
        if op == "XOR" and wire[0] not in ["x", "y", "z"] and left[0] not in ["x", "y", "z"] and right[0] not in ["x", "y", "z"]:
            wrong.add(wire)
        if op == "AND" and "x00" not in [left, right]:
            for w, (l, o, r) in gates.items():
                if (wire == l or wire == r) and o != "OR":
                    wrong.add(wire)
        if op == "XOR":
            for w, (l, o, r) in gates.items():
                if (wire == l or wire == r) and o == "OR":
                    wrong.add(wire)
    return wrong


def view(gates: dict[tuple[str, str, str, str], str]) -> None:
    circuit = Digraph()
    circuit.attr(rankdir="LR")
    added = set()
    count_op = {"AND": 0, "OR": 0, "XOR": 0}
    color_op = {"AND": "red", "OR": "green", "XOR": "blue"}
    for (left, op, right, _), wire in gates.items():
        circuit.node(op + str(count_op[op]), op, shape="box", color=color_op[op])
        count_op[op] += 1
        for node in [left, right, wire]:
            if node not in added:
                if node.startswith("x") or node.startswith("y") or node.startswith("z"):
                    circuit.node(node, node, shape="circle", color="yellow")
                else:
                    circuit.node(node, node, shape="ellipse")
                added.add(node)
            circuit.edge(node, op + str(count_op[op]))
    circuit.render("circuit", format="png", cleanup=True)
    circuit.view()


def part1():
    wires, gates = processing()
    total_wires, count = len(gates), 0
    operators = {"AND": lambda x, y: x & y, "OR": lambda x, y: x | y, "XOR": lambda x, y: x ^ y}
    while total_wires != count:
        for wire, (left, op, right) in gates.items():
            if wire not in wires and left in wires and right in wires:
                wires[wire] = operators[op](wires[left], wires[right])
                count += 1
    return int("".join(map(str, reversed([wires[wire] for wire in sorted(wires) if wire.startswith("z")]))), 2)


# Specific for n = 0 -> Half-adder
# Direct XOR        : x00 XOR y00 -> z00    => Gate H0
# Direct AND        : x00 AND y00 -> cout   => Gate H1
# General gates
# Entry XOR         : x.. XOR y.. -> val1   => Gate 0
# Intermediate XOR  : cin XOR val1 -> z..   => Gate 1
# Entry AND         : x.. AND y.. -> val2   => Gate 2
# Intermediate AND  : val1 AND cin -> val3  => Gate 3
# Exit OR           : val2 OR val3 -> cout  => Gate 4
# Specific for n = 45 -> Final bit
# Final OR          : val2 OR val3 -> z45   => Gate F0
def part2():
    return ",".join(sorted(checks(processing()[1])))


if __name__ == "__main__":
    print(part1())
    print(part2())
