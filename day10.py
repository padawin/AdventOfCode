numbers = list(range(256))
lengths = [225, 171, 131, 2, 35, 5, 0, 13, 1, 246, 54, 97, 255, 98, 254, 110]


def knot_hash(numbers, lengths):
    nb_numbers = len(numbers)
    position = 0
    skip = 0
    for length in lengths:
        # reverse order
        tmp = numbers + numbers
        start = position + length - 1
        end = (position - 1) if position is not 0 else None
        reversed_range = tmp[start:end:-1]
        for i, val in enumerate(reversed_range):
            numbers[(position + i) % nb_numbers] = val

        # update position
        position = (position + length + skip) % nb_numbers
        skip += 1

    return numbers


test_result = knot_hash([0, 1, 2, 3, 4], [3, 4, 1, 5])
assert test_result[0] == 3
assert test_result[1] == 4
res = knot_hash(numbers, lengths)
print(res[0] * res[1])
