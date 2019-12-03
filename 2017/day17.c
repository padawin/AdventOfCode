#include <stdio.h>

int main() {
	int steps = 382;
    int current_position = 0;
    int next_zero = 0;
    for (int i = 1; i < 50000000; ++i) {
        current_position = (current_position + steps) % i;
        if (current_position == 0) {
            next_zero = i;
		}
        current_position += 1;
	}

    printf("Val is: %d\n", next_zero);
	return 1;
}
