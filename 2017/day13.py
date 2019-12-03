layers = {
    0: 3,
    1: 2,
    2: 4,
    4: 4,
    6: 5,
    8: 6,
    10: 6,
    12: 8,
    14: 6,
    16: 6,
    18: 9,
    20: 8,
    22: 8,
    24: 8,
    26: 12,
    28: 8,
    30: 12,
    32: 12,
    34: 12,
    36: 10,
    38: 14,
    40: 12,
    42: 10,
    44: 8,
    46: 12,
    48: 14,
    50: 12,
    52: 14,
    54: 14,
    56: 14,
    58: 12,
    62: 14,
    64: 12,
    66: 12,
    68: 14,
    70: 14,
    72: 14,
    74: 17,
    76: 14,
    78: 18,
    84: 14,
    90: 20,
    92: 14,
}

test = {
    0: 3,
    1: 2,
    4: 4,
    6: 4
}


def get_severity(layers, shift):
    return [rank * depth
            for rank, depth in layers.items()
            if depth == 1 or (shift + rank) % ((depth - 1) * 2) == 0]


assert len(get_severity(test, 0)) > 0
assert sum(get_severity(test, 0)) == 24
assert len(get_severity(test, 10)) == 0
print(sum(get_severity(layers, 0)))


has_severity = True
wait = -1
while True:
    wait += 1
    try:
        first = next(rank * depth for rank, depth in layers.items() if depth == 1 or (wait + rank) % ((depth - 1) * 2) == 0)
    except StopIteration:
        break

print(wait)


# rank  > 0  1  2  3  4  5  6  7  8  9 10 11 12 13 14 15 16 17 18 19 20
# depth V
#     1   1  1  1  1  1  1  1  1  1  1  1  1  1  1  1  1  1  1  1  1  1
#     2   1  0  1  0  1  0  1  0  1  0  1  0  1  0  1  0  1  0  1  0  1
#     3   1  0  0  0  1  0  0  0  1  0  0  0  1  0  0  0  1  0  0  0  1
#     4   1  0  0  0  0  0  1  0  0  0  0  0  1  0  0  0  0  0  1  0  1
#     5   1  0  0  0  0  0  0  0  1  0  0  0  0  0  0  0  1  0  0  0  0
#
