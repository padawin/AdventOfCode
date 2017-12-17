steps = 382
steps_test = 3


def run_insert(steps):
    l = [0]
    current_position = 0
    for i in range(1, 2018):
        current_position = (current_position + steps) % len(l)
        l = l[:current_position+1] + [i] + l[current_position+1:]
        current_position += 1

    return l, current_position


l, current_position = run_insert(steps_test)
assert l[current_position + 1] == 638
l, current_position = run_insert(steps)
print(l[current_position + 1])
