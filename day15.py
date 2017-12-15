def create_gen(power, factor):
    mod = 2147483647
    while True:
        power = ((power % mod) * factor) % mod
        yield power


def find_nb_common(a, b):
    gen_a = create_gen(a, 16807)
    gen_b = create_gen(b, 48271)
    return sum(
        1
        for i in range(40000000)
        if next(gen_a) & 65535 == next(gen_b) & 65535
    )


test_a = 65
test_b = 8921
assert find_nb_common(test_a, test_b) == 588
print(find_nb_common(618, 814))
