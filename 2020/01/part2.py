import sys
import bisect


met = list()
metSet = set()
for line in sys.stdin:
    val = int(line)
    bisect.insort(met, val)
    metSet.add(val)

for a in met:
    b = 0
    while met[b] < 2020-a:
        if (2020-a-met[b]) in metSet:
            print(a*met[b]*(2020-a-met[b]))
            sys.exit(0)
        b += 1
