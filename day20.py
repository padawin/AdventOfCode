import re
import io
import sys


def find_closest_particle(particles):
    closest = None
    for index, particle in enumerate(particles):
        if (
            closest is None or
            closest['a'] > particle['a'] or
            closest['a'] == particle['a'] and closest['s'] > particle['s'] or
            closest['a'] == particle['a'] and closest['s'] == particle['s'] and closest['p'] > particle['p']
        ):
            closest = {'index': index, **particle}

    return closest


def get_particles(stream):
    particles = []
    reg = 'p=<([^,]+),([^,]+),([^>]+)>, v=<([^,]+),([^,]+),([^>]+)>, a=<([^,]+),([^,]+),([^>]+)>'
    for line in stream.readlines():
        p = re.search(reg, line.rstrip('\n'))
        particle = {
            'p': abs(int(p.group(1))) + abs(int(p.group(2))) + abs(int(p.group(3))),
            's': abs(int(p.group(4))) + abs(int(p.group(5))) + abs(int(p.group(6))),
            'a': abs(int(p.group(7))) + abs(int(p.group(8))) + abs(int(p.group(9)))
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
