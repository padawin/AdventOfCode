import io
import sys


def create_graph(input_data):
    # a vertice is a number
    # an edge is a couple of numbers, a token. Each line of the input is an edge
    edges = list()
    edges_per_vertex = dict()
    for edge_index, line in enumerate(input_data.readlines()):
        edge = line.rstrip('\n')
        vertex1, vertex2 = (int(vertex) for vertex in edge.split('/'))
        edges.append((vertex1, vertex2))
        if vertex1 not in edges_per_vertex:
            edges_per_vertex[vertex1] = set()
        if vertex2 not in edges_per_vertex:
            edges_per_vertex[vertex2] = set()
        edges_per_vertex[vertex1].add(edge_index)
        edges_per_vertex[vertex2].add(edge_index)

    return edges, edges_per_vertex


def _find_path(edges_per_vertex, branches):
    current_vertex, edges_to_process, current_path = branches.pop(0)
    while not edges_to_process.issubset(current_path):
        next_edge = edges_to_process.pop()
        if len(edges_to_process) > 0:
            branches.append((current_vertex, edges_to_process, current_path.copy()))
        current_path.add(next_edge)
        candidate_vertex = list(edges[next_edge])
        candidate_vertex.remove(current_vertex)
        current_vertex = candidate_vertex.pop()
        edges_to_process = edges_per_vertex[current_vertex] - current_path
    return current_path


def find_strongest_path(edges, edges_per_vertex):
    max_strength = 0
    max_path = None
    current_path = set()
    current_vertex = 0
    branches = [(current_vertex, edges_per_vertex[current_vertex], current_path)]
    loop = 0
    while len(branches):
        path = _find_path(edges_per_vertex, branches)
        strength_path = _calculate_path_strength(edges, path)
        if strength_path > max_strength:
            max_strength = strength_path
            max_path = path

    return max_path, max_strength


def find_longest_path(edges, edges_per_vertex):
    max_length = 0
    max_strength = 0
    max_path = None
    current_path = set()
    current_vertex = 0
    branches = [(current_vertex, edges_per_vertex[current_vertex], current_path)]
    loop = 0
    while len(branches):
        path = _find_path(edges_per_vertex, branches)
        length_path = len(path)
        # print(length_path)
        strength_path = _calculate_path_strength(edges, path)
        if (
            length_path == max_length and strength_path > max_strength
            or length_path > max_length
        ):
            max_strength = strength_path
            max_length = length_path
            max_path = path

    return max_path, max_length, max_strength


def _calculate_path_strength(edges, path):
    return sum(sum(edges[edge]) for edge in path)


test_data = """\
0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10
"""
edges, edges_per_vertex = create_graph(io.StringIO(test_data))
path, strength = find_strongest_path(edges, edges_per_vertex)
assert strength == 31


edges, edges_per_vertex = create_graph(sys.stdin)
path, length, strength = find_longest_path(edges, edges_per_vertex)
print(strength)
