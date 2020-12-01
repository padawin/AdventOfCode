import sys

met = dict()

for line in sys.stdin:
    val = int(line)
    if val in met:
        print(val * met[val])
        break

    met[2020-val] = val
