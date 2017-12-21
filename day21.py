import io
import sys
import math


def get_nb_active_bits(s):
    s = ''.join(s)
    active = 0
    for c in s:
        if c == '1':
            active += 1
    return active


def _rotate_clockwise(matrix):
    return [''.join(line) for line in zip(*matrix[::-1])]


def _analyse_rules(rules):
    res = {2: {}, 3: {}}
    for line in rules.readlines():
        template, _, result = line.rstrip('\n').replace('.', '0').replace('#', '1').split(' ')
        res[int(math.sqrt(len(template)))][template] = result.split('/')
    return res


def _divide_matrix(matrix):
    size = len(matrix)
    if size % 2 == 0:
        step = 2
    elif size % 3 == 0:
        step = 3
    else:
        raise ValueError()

    submatrices = []
    for j in range(0, size, step):
        for i in range(0, size, step):
            sub = [matrix[row][i:i+step] for row in range(j, j + step)]
            submatrices.append(sub)
    return submatrices, step


def _find_rule(matrix, rules):
    for rule, result in rules.items():
        if (
            _matrices_are_similar(matrix, rule.split('/')) or
            _matrices_are_mirror(matrix, rule.split('/'))
        ):
            return result
    return matrix


def _matrices_are_similar(matrix1, matrix2):
    orig = matrix1
    while matrix1 != matrix2:
        matrix1 = _rotate_clockwise(matrix1)
        if matrix1 == orig:
            break
    return matrix1 == matrix2


def _matrices_are_mirror(matrix1, matrix2):
    matrix1 = matrix1[::-1]
    orig = matrix1
    while matrix1 != matrix2:
        matrix1 = _rotate_clockwise(matrix1)
        if matrix1 == orig:
            break
    return matrix1 == matrix2


def _recreate_matrix(submatrices):
    shift = 0
    matrix = []
    nb_submatrices = len(submatrices)
    if nb_submatrices == 1:
        return submatrices[0]
    width = int(math.sqrt(nb_submatrices))
    while shift < nb_submatrices:
        matrix += [''.join(l) for l in zip(*submatrices[shift:shift+width])]
        shift += width
    return matrix


def process(rules, steps):
    matrix = [
        '010',
        '001',
        '111'
    ]

    rules = _analyse_rules(rules)
    for s in range(steps):
        tmp = []
        submatrices, size = _divide_matrix(matrix)
        for submatrix in submatrices:
            tmp.append(_find_rule(submatrix, rules[size]))
        matrix = _recreate_matrix(tmp)
    return matrix


test_input = """\
../.# => ##./#../...
.#./..#/### => #..#/..../..../#..#
"""

new_matrix = process(io.StringIO(test_input), 2)
assert new_matrix == [
    '110110',
    '100100',
    '000000',
    '110110',
    '100100',
    '000000'
]
assert get_nb_active_bits(new_matrix) == 12

new_matrix = process(sys.stdin, 18)
print(get_nb_active_bits(new_matrix))
