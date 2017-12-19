import io
import sys


def build_grid(input_data):
    grid = []
    read = True
    width = 0
    start_col = None
    rows = [line.rstrip('\n') for line in input_data.readlines()]
    for row_str in rows:
        row = []
        for col, c in enumerate(row_str):
            row.append(c)
            if start_col is None and c == '|':
                start_col = col
        grid.append(row)
        if width == 0:
            width = len(row)
    return grid, start_col


def _change_direction(grid, curr, direction):
    if direction in ('left', 'right'):
        try:
            direction = 'up' if grid[curr['y'] - 1][curr['x']] != ' ' else 'down'
        except IndexError:
            direction = 'down'
    elif direction in ('up', 'down'):
        try:
            direction = 'left' if grid[curr['y']][curr['x'] - 1] != ' ' else 'right'
        except IndexError:
            direction = 'right'
    return direction


def travel(grid, start_x):
    direction = 'down'
    curr = {'x': start_x, 'y': 0}
    letters = []
    steps = 0
    end_reached = False
    while not end_reached:
        curr_char = grid[curr['y']][curr['x']]
        if curr_char == ' ':
            break
        if curr_char == '+':
            direction = _change_direction(grid, curr, direction)
        elif curr_char not in ('-', '|'):
            letters.append(curr_char)

        if direction == 'down':
            curr['y'] += 1
        elif direction == 'up':
            curr['y'] -= 1
        elif direction == 'left':
            curr['x'] -= 1
        elif direction == 'right':
            curr['x'] += 1
        steps += 1

    return letters, steps



test_data = [
    '     |          ',
    '     |  +--+    ',
    '     A  |  C    ',
    ' F---|----E|--+ ',
    '     |  |  |  D ',
    '     +B-+  +--+ '
]

grid, start_col = build_grid(io.StringIO('\n'.join(test_data)))
assert start_col == 5
assert grid == [
    [' ', ' ', ' ', ' ', ' ', '|', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '],
    [' ', ' ', ' ', ' ', ' ', '|', ' ', ' ', '+', '-', '-', '+', ' ', ' ', ' ', ' '],
    [' ', ' ', ' ', ' ', ' ', 'A', ' ', ' ', '|', ' ', ' ', 'C', ' ', ' ', ' ', ' '],
    [' ', 'F', '-', '-', '-', '|', '-', '-', '-', '-', 'E', '|', '-', '-', '+', ' '],
    [' ', ' ', ' ', ' ', ' ', '|', ' ', ' ', '|', ' ', ' ', '|', ' ', ' ', 'D', ' '],
    [' ', ' ', ' ', ' ', ' ', '+', 'B', '-', '+', ' ', ' ', '+', '-', '-', '+', ' ']
]

letters, steps = travel(grid, start_col)
assert ''.join(letters) == 'ABCDEF'
assert steps == 38

grid, start_col = build_grid(sys.stdin)
letters, steps = travel(grid, start_col)
print(steps)
