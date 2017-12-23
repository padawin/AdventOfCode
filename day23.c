#include <stdio.h>

int main() {
	int from = 108400;
	int to = 125400;
	int n = 0;
	int nbNonPrime = 0;

	for (; from <= to; from += 17) {
		for (n = 2; n < from; ++n) {
			if ((from % n) == 0) {
				nbNonPrime++;
				break;
			}
		}
	}
	printf("%d\n", nbNonPrime);
	return 0;
}
