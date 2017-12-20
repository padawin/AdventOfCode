import re
import io
import sys


def find_closest_particle(particles):
    closest = None
    for index, particle in enumerate(particles):
        a = sum(abs(val) for val in particle['a'])
        s = sum(abs(val) for val in particle['s'])
        p = sum(abs(val) for val in particle['p'])
        if (
            closest is None or
            closest['a'] > a or
            closest['a'] == a and closest['s'] > s or
            closest['a'] == a and closest['s'] == s and closest['p'] > p
        ):
            closest = {'index': index, 'a': a, 's': s, 'p': p}

    return closest


def get_particles(stream):
    particles = []
    reg = 'p=<([^,]+),([^,]+),([^>]+)>, v=<([^,]+),([^,]+),([^>]+)>, a=<([^,]+),([^,]+),([^>]+)>'
    for line in stream.readlines():
        p = re.search(reg, line.rstrip('\n'))
        particle = {
            'p': (int(p.group(1)), int(p.group(2)), int(p.group(3))),
            's': (int(p.group(4)), int(p.group(5)), int(p.group(6))),
            'a': (int(p.group(7)), int(p.group(8)), int(p.group(9)))
        }
        particles.append(particle)
    return particles


test_data_1 = """\
p=< 3,0,0>, v=< 2,0,0>, a=<-1,0,0>
p=< 4,0,0>, v=< 0,0,0>, a=<-2,0,0>
"""
test_data_2 = """\
p=< 3,0,0>, v=< 2,0,0>, a=<-2,0,0>
p=< 4,0,0>, v=< 0,0,0>, a=<-2,0,0>
"""
test_data_3 = """\
p=< 3,0,0>, v=< 2,0,0>, a=<-1,0,0>
p=< 4,0,0>, v=< 2,0,0>, a=<-2,0,0>
"""
particles_1 = get_particles(io.StringIO(test_data_1))
assert find_closest_particle(particles_1)['index'] == 0
particles_2 = get_particles(io.StringIO(test_data_2))
assert find_closest_particle(particles_2)['index'] == 1
particles_3 = get_particles(io.StringIO(test_data_3))
assert find_closest_particle(particles_3)['index'] == 0

particles = get_particles(sys.stdin)
print(find_closest_particle(particles)['index'])
