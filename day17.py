steps = 382
steps_test = 3


def run_insert(steps):
    current_position = 0
    next_zero = None
    for i in range(1, 50000000):
        current_position = (current_position + steps) % i
        if current_position == 0:
            next_zero = i
        if i == 2017:
            assert next_zero == 41
        current_position += 1

    return next_zero


next_zero = run_insert(steps)
print(next_zero)
