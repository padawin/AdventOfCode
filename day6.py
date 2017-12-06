input_vals = [10, 3, 15, 10, 5, 15, 5, 15, 9, 2, 5, 8, 5, 2, 3, 6]


def sign(vals):
    return '-'.join(str(i) for i in vals)


nb_elements = len(input_vals)
states = set()
states.add(sign(input_vals))

nb_iterations = 0
while True:
    max_val = max(input_vals)
    index_val = input_vals.index(max_val)
    input_vals[index_val] = 0
    for i in range(max_val):
        input_vals[(index_val + 1 + i) % nb_elements] += 1
    nb_states = len(states)
    states.add(sign(input_vals))
    nb_iterations += 1
    if nb_states == len(states):
        break

print(nb_iterations)
