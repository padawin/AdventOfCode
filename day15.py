from math import pow


def find_nb_common(a, b):
    factor_a = 16807
    factor_b = 48271
    mod = 2147483647
    previous_power_a = 1
    previous_power_b = 1
    nb = 0
    for i in range(1, 40000001):
        power_a = (previous_power_a % mod) * factor_a
        previous_power_a = power_a
        power_b = (previous_power_b % mod) * factor_b
        previous_power_b = power_b
        calc_a = (a * (power_a % mod)) % mod
        calc_b = (b * (power_b % mod)) % mod
        if calc_a & 65535 == calc_b & 65535:
            nb += 1
    return nb


test_a = 65
test_b = 8921
# assert find_nb_common(test_a, test_b) == 588
print(find_nb_common(618, 814))
