import io
import sys


def _get_computer_key(x, y):
    return '{}-{}'.format(x, y)


def analyse_grid(grid_in):
    grid = [l.strip('\n') for l in grid_in.readlines()]
    infected_computers = set(
        _get_computer_key(x, y)
        for y, row in enumerate(grid)
        for x, val in enumerate(row)
        if val == '#'
    )
    grid_center = [int(len(grid[0]) / 2), int(len(grid) / 2)]
    return infected_computers, grid_center


def update_virus(computer_infected, virus):
    # up right down left
    directions = ((0, -1), (1, 0), (0, 1), (-1, 0))
    nb_directions = len(directions)
    if computer_infected:
        way = 1
    else:
        way = -1

    virus['orientation'] = (nb_directions + virus['orientation'] + way) % nb_directions
    virus['position'] = [
        sum(coord)
        for coord in zip(
            virus['position'],
            directions[virus['orientation']]
        )
    ]


def propagate(virus_position, initial_state, steps):
    virus = {'position': [*virus_position], 'orientation': 0}
    step = 0
    nb_infections = 0
    infected_computers = initial_state.copy()
    while step < steps:
        computer_key = _get_computer_key(*virus['position'])
        computer_infected = computer_key in infected_computers
        update_virus(computer_infected, virus)
        if computer_infected:
            infected_computers.remove(computer_key)
        else:
            infected_computers.add(computer_key)
            nb_infections += 1
        step += 1

    return nb_infections


test_input = """\
..#
#..
...
"""

infected_computers, grid_center = analyse_grid(io.StringIO(test_input))
nb_infections = propagate(grid_center, infected_computers, 70)
assert nb_infections == 41
nb_infections = propagate(grid_center, infected_computers, 10000)
assert nb_infections == 5587

infected_computers, grid_center = analyse_grid(sys.stdin)
nb_infections = propagate(grid_center, infected_computers, 10000)
print(nb_infections)
