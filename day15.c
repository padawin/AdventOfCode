#include <stdio.h>

int find_nb_common(int a, int b) {
    int factor_a = 16807;
    int factor_b = 48271;
    long mod = 2147483647;
    long previous_power_a = 1;
    long previous_power_b = 1;
    int nb = 0;
	for (int i = 0; i < 40000000; ++i) {
        long power_a = (previous_power_a % mod) * factor_a;
        previous_power_a = power_a;
        long power_b = (previous_power_b % mod) * factor_b;
        previous_power_b = power_b;
        long calc_a = (a * (power_a % mod)) % mod;
        long calc_b = (b * (power_b % mod)) % mod;
        if ((calc_a & 65535) == (calc_b & 65535)) {
            nb += 1;
		}
	}
    return nb;
}

int main() {
	printf("Nb is: %d\n", find_nb_common(618, 814));
	return 1;
}
