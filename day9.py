import sys


def process_line(l):
    score = 0
    level = 0
    for c in l:
        if c == '{':
            level += 1
            score += level
        elif c == '}':
            level -= 1
    return score


for line in sys.stdin:
    l = line.strip()
    res = process_line(l)
    print("For line {}, the score is: {}".format(l, res))
