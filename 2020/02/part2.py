import sys
import re

countValid = 0
for line in sys.stdin:
    p = re.compile(r'(\d+)-(\d+) ([a-z]): ([a-z]+)')
    match = p.match(line)
    pos1 = int(match.group(1))-1
    pos2 = int(match.group(2))-1
    letter = match.group(3)
    password = match.group(4)
    if password[pos1] == letter and password[pos2] != letter or password[pos2] == letter and password[pos1] != letter:
        countValid += 1

print(countValid)
