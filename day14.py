import day10

test = 'flqrgnkx'


def get_nb_active_bits(s):
    grid = []
    for i in range(128):
        grid.append(day10.knot('{}-{}'.format(s, i)))
    grid = int('0x{}'.format(''.join(grid)), 0)
    c = 0
    while grid:
        c += 1
        grid &= grid - 1

    return c


assert get_nb_active_bits(test) == 8108
print(get_nb_active_bits('ffayrhll'))
