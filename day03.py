import math

# 37 36 35 34 33 32 31
# 38 17 16 15 14 13 30
# 39 18  5  4  3 12 29
# 40 19  6  1  2 11 28
# 41 20  7  8  9 10 27
# 42 21 22 23 24 25 26
# 43 44 45 46 47 48 49
#
# 1 0 0
# 2 1 0
# 3
# 4 0 1
fixtures = [
    [0, 0, 0],
    [1, 0, 0],
    [2, 1, 0],
    [3, 1, -1],
    [4, 0, -1],
    [5, -1, -1],
    [6, -1, 0],
    [7, -1, 1],
    [8, 0, 1],
    [9, 1, 1],
    [10, 2, 1],
    [11, 2, 0],
    [12, 2, -1],
    [13, 2, -2],
    [14, 1, -2],
    [15, 0, -2],
    [16, -1, -2],
    [17, -2, -2],
    [18, -2, -1],
    [19, -2, 0],
    [20, -2, 1],
    [21, -2, 2],
    [22, -1, 2],
    [23, 0, 2],
    [24, 1, 2],
    [25, 2, 2],
    [26, 3, 2],
    [27, 3, 1],
    [28, 3, 0],
    [29, 3, -1],
    [30, 3, -2],
    [31, 3, -3],
    [32, 2, -3],
    [33, 1, -3],
    [34, 0, -3],
    [35, -1, -3],
    [36, -2, -3],
    [37, -3, -3],
    [38, -3, -2],
    [39, -3, -1],
    [40, -3, 0],
    [41, -3, 1],
    [42, -3, 2],
    [43, -3, 3],
    [44, -2, 3],
    [45, -1, 3],
    [46, 0, 3],
    [47, 1, 3],
    [48, 2, 3],
    [49, 3, 3]
]


def get_x(n, sign):
    shift = -1 if n % 2 == 0 else 0
    return int(sign * int(math.sqrt(n) / 2)) - shift


def get_y(n, sign):
    return int(sign * int(math.sqrt(n) / 2))


def get_sign(n):
    return -1 if n % 2 == 0 else 1


def coords(n):
    sqrt = int(math.sqrt(n))
    closest_sq = math.pow(sqrt, 2)
    sign = get_sign(closest_sq)
    diff = n - closest_sq
    val_closest_sq = (get_x(closest_sq, sign), get_y(closest_sq, sign))
    if diff == 0:
        return val_closest_sq
    elif diff <= sqrt + 1:
        return (val_closest_sq[0] + sign, val_closest_sq[1] - sign * (diff - 1))
    else:
        x = n - (closest_sq + sqrt + 2)
        return (val_closest_sq[0] - (sign * x), -1 * val_closest_sq[1] - max(0, sign))


for n in range(1, len(fixtures) - 1):
    c = coords(n)
    assert c[0] == fixtures[n][1]
    assert c[1] == fixtures[n][2]

print(sum(abs(coord) for coord in coords(368078)))
