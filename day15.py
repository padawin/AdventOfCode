def find_nb_common(a, b):
    factor_a = 16807
    factor_b = 48271
    mod = 2147483647
    power_a = a
    power_b = b
    nb = 0
    for i in range(40000000):
        power_a = ((power_a % mod) * factor_a) % mod
        power_b = ((power_b % mod) * factor_b) % mod
        if power_a & 65535 == power_b & 65535:
            nb += 1
    return nb


test_a = 65
test_b = 8921
assert find_nb_common(test_a, test_b) == 588
print(find_nb_common(618, 814))
