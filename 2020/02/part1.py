import sys
import re

countValid = 0
for line in sys.stdin:
    p = re.compile(r'(\d+)-(\d+) ([a-z]): ([a-z]+)')
    match = p.match(line)
    minCount = int(match.group(1))
    maxCount = int(match.group(2))
    letter = match.group(3)
    password = match.group(4)
    if minCount <= (len(password) - len(password.replace(letter, ''))) <= maxCount:
        countValid += 1
print(countValid)
