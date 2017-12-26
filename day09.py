import sys


def process_line(l):
    print(len(l) - 2)
    return len(l) - 2


print(sum(process_line(l.strip()) for l in sys.stdin))
