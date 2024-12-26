from networkx import Graph, enumerate_all_cliques, find_cliques


def processing() -> Graph:
    connections = [tuple(connection.split("-")) for connection in open("src/day23/input.txt").read().splitlines()]
    graph = Graph()
    graph.add_nodes_from(connection[0] for connection in connections)
    graph.add_edges_from(connections)
    return graph


def part1() -> int:
    return len(list(filter(lambda clique: len(clique) == 3 and any(computer.startswith("t") for computer in clique),
                           enumerate_all_cliques(processing()))))


def part2() -> str:
    return ",".join(sorted(max(list(find_cliques(processing())), key=len)))


if __name__ == "__main__":
    print(part1())
    print(part2())
