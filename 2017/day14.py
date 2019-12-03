import day10

test = 'flqrgnkx'


def get_grid_int(s):
    grid = []
    for i in range(128):
        grid.append(day10.knot('{}-{}'.format(s, i)))
    return int('0x{}'.format(''.join(grid)), 0)


def get_nb_active_bits(s):
    grid = get_grid_int(s)
    c = 0
    while grid:
        c += 1
        grid &= grid - 1

    return c


def get_regions(grid):
    nb_regions = 0
    i = 0
    while grid:
        if grid & 1 << i:
            x = i % 128
            y = i // 128
            grid = _get_region(grid, x, y)
            nb_regions += 1
        i += 1
    return nb_regions


def _get_region(grid, x, y):
    to_visit = {(x, y)}
    visited = {}
    while to_visit:
        current = to_visit.pop()
        i = current[1] * 128 + current[0]
        if grid & 1 << i:
            grid ^= 1 << i
            if current[0] > 0:
                to_visit.add((current[0] - 1, current[1]))
            if current[0] < 127:
                to_visit.add((current[0] + 1, current[1]))
            if current[1] > 0:
                to_visit.add((current[0], current[1] - 1))
            if current[1] < 127:
                to_visit.add((current[0], current[1] + 1))
    return grid


def _dump_grid(grid):
    print("Dump:")
    grid_str = bin(grid)[2:]
    [print(grid_str[i:i+128]) for i in range(0, len(grid_str), 128)]


assert get_nb_active_bits(test) == 8108
assert get_nb_active_bits('ffayrhll') == 8190

assert get_regions(get_grid_int(test)) == 1242
print(get_regions(get_grid_int('ffayrhll')))
