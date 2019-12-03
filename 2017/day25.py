import io
import sys
import textwrap
import re


def translate_blueprint(blueprint):
    output_pattern = textwrap.dedent("""\
    state = '{}'
    steps = {}

    actions = {{
    {}
    }}

    position = 0
    active_cells = dict()
    for step in range(steps):
        active_cells[position], move, state = actions[state](0 if position not in active_cells else active_cells[position])
        position += move

    print(sum(active_cells.values()))
    """)
    lines = [line.rstrip('\n') for line in blueprint]
    initial_state_regex = '^Begin in state ([A-Z]).$'
    steps_number_regex = '^Perform a diagnostic checksum after ([0-9]+) steps.$'
    state_program_regex = '^In state ([A-Z]):$'
    set_val_regex = '^    - Write the value (0|1).$'
    move_direction_regex = '^    - Move one slot to the (left|right).$'
    new_state_regex = '^    - Continue with state ([A-Z]).$'

    res = re.search(initial_state_regex, lines[0])
    initial_state = res.group(1)
    res = re.search(steps_number_regex, lines[1])
    nb_steps = res.group(1)
    states = []
    state_pattern = textwrap.dedent("""'{state_from}': lambda val: ({val_0}, {move_0}, '{state_to_0}') if val == 0 else ({val_1}, {move_1}, '{state_to_1}'),""")
    for line in range(2, len(lines), 10):
        res = re.search(state_program_regex, lines[line+1])
        state_from = res.group(1)
        res = re.search(set_val_regex, lines[line+3])
        val_0 = res.group(1)
        res = re.search(move_direction_regex, lines[line+4])
        move_0 = 1 if res.group(1) == 'right' else -1
        res = re.search(new_state_regex, lines[line+5])
        state_to_0 = res.group(1)
        res = re.search(set_val_regex, lines[line+7])
        val_1 = res.group(1)
        res = re.search(move_direction_regex, lines[line+8])
        move_1 = 1 if res.group(1) == 'right' else -1
        res = re.search(new_state_regex, lines[line+9])
        state_to_1 = res.group(1)
        states.append(state_pattern.format(
            state_from=state_from,
            val_0=val_0,
            move_0=move_0,
            state_to_0=state_to_0,
            val_1=val_1,
            move_1=move_1,
            state_to_1=state_to_1
        ))

    return output_pattern.format(initial_state, nb_steps, ''.join(states))


test_input = """\
Begin in state A.
Perform a diagnostic checksum after 6 steps.

In state A:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state B.
  If the current value is 1:
    - Write the value 0.
    - Move one slot to the left.
    - Continue with state B.

In state B:
  If the current value is 0:
    - Write the value 1.
    - Move one slot to the left.
    - Continue with state A.
  If the current value is 1:
    - Write the value 1.
    - Move one slot to the right.
    - Continue with state A.
"""
test_output = """\
state = 'A'
steps = 6

actions = {
'A': lambda val: (1, 1, 'B') if val == 0 else (0, -1, 'B'),'B': lambda val: (1, -1, 'A') if val == 0 else (1, 1, 'A'),
}

position = 0
active_cells = dict()
for step in range(steps):
    active_cells[position], move, state = actions[state](0 if position not in active_cells else active_cells[position])
    position += move

print(sum(active_cells.values()))
"""

output = translate_blueprint(io.StringIO(test_input))
assert test_output == output

output = translate_blueprint(sys.stdin)
print(output)
