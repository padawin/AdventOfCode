import io
import sys


STATE_WEAKENED = 'W'
STATE_INFECTED = 'I'
STATE_FLAGGED = 'F'


def _get_computer_key(x, y):
    return '{}-{}'.format(x, y)


def analyse_grid(grid_in):
    grid = [l.strip('\n') for l in grid_in.readlines()]
    infected_computers = {
        _get_computer_key(x, y): STATE_INFECTED
        for y, row in enumerate(grid)
        for x, val in enumerate(row)
        if val == '#'
    }
    grid_center = [int(len(grid[0]) / 2), int(len(grid) / 2)]
    return infected_computers, grid_center


def update_virus(infected_computers, computer, virus):
    # up right down left
    directions = ((0, -1), (1, 0), (0, 1), (-1, 0))
    nb_directions = len(directions)

    states_directions = {
        STATE_WEAKENED: 0,
        STATE_INFECTED: 1,
        STATE_FLAGGED: 2
    }
    try:
        way = states_directions[infected_computers[computer]]
    except KeyError:
        way = -1

    virus['orientation'] = (nb_directions + virus['orientation'] + way) % nb_directions
    virus['position'] = [
        sum(coord)
        for coord in zip(
            virus['position'],
            directions[virus['orientation']]
        )
    ]


def update_computer_state(infected_computers, computer):
    states = {
        STATE_WEAKENED: STATE_INFECTED,
        STATE_INFECTED: STATE_FLAGGED,
        STATE_FLAGGED: None
    }
    try:
        infected_computers[computer] = states[infected_computers[computer]]
    except KeyError:
        infected_computers[computer] = STATE_WEAKENED
    else:
        if infected_computers[computer] is None:
            del infected_computers[computer]
        elif infected_computers[computer] == STATE_INFECTED:
            return True

    return False


def propagate(virus_position, initial_state, steps):
    virus = {'position': [*virus_position], 'orientation': 0}
    step = 0
    nb_infections = 0
    infected_computers = initial_state.copy()
    while step < steps:
        computer_key = _get_computer_key(*virus['position'])
        update_virus(infected_computers, computer_key, virus)
        infected = update_computer_state(infected_computers, computer_key)
        step += 1
        if infected:
            nb_infections += 1

    return nb_infections


test_input = """\
..#
#..
...
"""

infected_computers, grid_center = analyse_grid(io.StringIO(test_input))
nb_infections = propagate(grid_center, infected_computers, 100)
assert nb_infections == 26
nb_infections = propagate(grid_center, infected_computers, 10000000)
assert nb_infections == 2511944

infected_computers, grid_center = analyse_grid(sys.stdin)
nb_infections = propagate(grid_center, infected_computers, 10000000)
print(nb_infections)
