#include <stdio.h>

int find_nb_common(int a, int b) {
    int factor_a = 16807;
    int factor_b = 48271;
    long mod = 2147483647;
    long power_a = a;
    long power_b = b;
    int nb = 0;
	for (int i = 0; i < 5000000; ++i) {
		do {
			power_a = ((power_a % mod) * factor_a) % mod;
		} while ((power_a % 4) != 0);

		do {
			power_b = ((power_b % mod) * factor_b) % mod;
		} while ((power_b % 8) != 0);
        if ((power_a & 65535) == (power_b & 65535)) {
            nb += 1;
		}
	}
    return nb;
}

int main() {
	printf("Nb is: %d\n", find_nb_common(618, 814));
	return 1;
}
