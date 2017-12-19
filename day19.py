def build_grid(input_data=None):
    grid = []
    read = True
    width = 0
    start_col = None
    while read:
        if input_data is not None:
            line = input_data.pop(0)
        else:
            line = input()
        row = []
        for col, c in enumerate(line):
            row.append(c)
            if start_col is None and c == '|':
                start_col = col
        grid.append(row)
        if width == 0:
            width = len(row)

        read = (
            isinstance(input_data, list) and len(input_data) != 0
            or input_data is None and line.strip() != ''
        )
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

    return letters



test_data = [
    '     |          ',
    '     |  +--+    ',
    '     A  |  C    ',
    ' F---|----E|--+ ',
    '     |  |  |  D ',
    '     +B-+  +--+ '
]

grid, start_col = build_grid(test_data)
assert start_col == 5
assert grid == [
    [' ', ' ', ' ', ' ', ' ', '|', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' '],
    [' ', ' ', ' ', ' ', ' ', '|', ' ', ' ', '+', '-', '-', '+', ' ', ' ', ' ', ' '],
    [' ', ' ', ' ', ' ', ' ', 'A', ' ', ' ', '|', ' ', ' ', 'C', ' ', ' ', ' ', ' '],
    [' ', 'F', '-', '-', '-', '|', '-', '-', '-', '-', 'E', '|', '-', '-', '+', ' '],
    [' ', ' ', ' ', ' ', ' ', '|', ' ', ' ', '|', ' ', ' ', '|', ' ', ' ', 'D', ' '],
    [' ', ' ', ' ', ' ', ' ', '+', 'B', '-', '+', ' ', ' ', '+', '-', '-', '+', ' ']
]

letters = travel(grid, start_col)
assert ''.join(letters) == 'ABCDEF'

grid, start_col = build_grid()
print(''.join(travel(grid, start_col)))
