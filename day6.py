input_vals = [10, 3, 15, 10, 5, 15, 5, 15, 9, 2, 5, 8, 5, 2, 3, 6]


def sign(vals):
    return '-'.join(str(i) for i in vals)


nb_elements = len(input_vals)
states = {}

nb_iterations = 0
states[sign(input_vals)] = nb_iterations
while True:
    max_val = max(input_vals)
    index_val = input_vals.index(max_val)
    input_vals[index_val] = 0
    for i in range(max_val):
        input_vals[(index_val + 1 + i) % nb_elements] += 1
    nb_states = len(states)
    nb_iterations += 1
    signed = sign(input_vals)
    if signed in states:
        break
    else:
        states[signed] = nb_iterations


print(nb_iterations - states[sign(input_vals)])
