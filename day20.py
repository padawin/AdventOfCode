import re
import math
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
        particle = (
            (int(p.group(1)), int(p.group(2)), int(p.group(3))),
            (int(p.group(4)), int(p.group(5)), int(p.group(6))),
            (int(p.group(7)), int(p.group(8)), int(p.group(9)))
        )
        particles.append(particle)
    return particles


def solve_equation(a, b, c):
    t = None
    if a != 0:
        d = (b ** 2) - 4 * a * c
        if d >= 0:
            t = max(
                (-b - math.sqrt(d)) / (2 * a),
                (-b + math.sqrt(d)) / (2 * a)
            )
    elif b != 0:
        t = -c / b
        t = t if t >= 0 else None
    elif c == 0:
        t = c

    return t


def remove_collisions(particles):
    nb_particles = len(particles)
    particles_indexes = set(range(len(particles)))
    collisions = {}
    for i, p1 in enumerate(particles):
        for j in range(i + 1, nb_particles):
            p2 = particles[j]
            ax, bx, cx = [p1[x][0] - p2[x][0] for x in ['a', 's', 'p']]
            ay, by, cy = [p1[y][1] - p2[y][1] for y in ['a', 's', 'p']]
            az, bz, cz = [p1[z][2] - p2[z][2] for z in ['a', 's', 'p']]
            collision = {
                c
                for c in [
                    solve_equation(ax, bx, cx),
                    solve_equation(ay, by, cy),
                    solve_equation(az, bz, cz)
                ]
                if c is not None
            }

            if len(collision) > 1 and 0 in collision:
                collision -= {0}
            if len(collision) == 1:
                collision = str(collision.pop())
                if collision in collisions:
                    collisions[collision].add(i)
                    collisions[collision].add(j)
                else:
                    collisions[collision] = {i, j}

    for collision_time in sorted(collisions, key=lambda collision: collision[0]):
        particles_colliding = collisions[collision_time]
        if particles_colliding.issubset(particles_indexes):
            particles_indexes -= particles_colliding

    return particles_indexes


# test_data_1 = """\
# p=< 3,0,0>, v=< 2,0,0>, a=<-1,0,0>
# p=< 4,0,0>, v=< 0,0,0>, a=<-2,0,0>
# """
# test_data_2 = """\
# p=< 3,0,0>, v=< 2,0,0>, a=<-2,0,0>
# p=< 4,0,0>, v=< 0,0,0>, a=<-2,0,0>
# """
# test_data_3 = """\
# p=< 3,0,0>, v=< 2,0,0>, a=<-1,0,0>
# p=< 4,0,0>, v=< 2,0,0>, a=<-2,0,0>
# """
# particles_1 = get_particles(io.StringIO(test_data_1))
# assert find_closest_particle(particles_1)['index'] == 0
# particles_2 = get_particles(io.StringIO(test_data_2))
# assert find_closest_particle(particles_2)['index'] == 1
# particles_3 = get_particles(io.StringIO(test_data_3))
# assert find_closest_particle(particles_3)['index'] == 0
#
# test_data_collisions = """\
# p=<-6,0,0>, v=< 3,0,0>, a=< 0,0,0>
# p=<-4,0,0>, v=< 2,0,0>, a=< 0,0,0>
# p=<-2,0,0>, v=< 1,0,0>, a=< 0,0,0>
# p=< 3,0,0>, v=<-1,0,0>, a=< 0,0,0>
# """
# particles = get_particles(io.StringIO(test_data_collisions))
# assert remove_collisions(particles) == {3}
#
# test_data_collisions = """\
# p=<-6,6,0>, v=< 3,-3,0>, a=< 0,0,0>
# p=<-4,0,0>, v=< 2,0,0>, a=< 0,0,0>
# p=<0,-2,0>, v=< 0,1,0>, a=< 0,0,0>
# p=< 3,0,0>, v=<-1,0,0>, a=< 0,0,0>
# """
# particles = get_particles(io.StringIO(test_data_collisions))
# assert remove_collisions(particles) == {3}
#
particles = get_particles(sys.stdin)
# print(find_closest_particle(particles)['index'])
# print(len(remove_collisions(particles)))

from collections import Counter

def run_p(p):
    (x, y, z), (vx, vy, vz), (ax, ay, az) = p
    vx, vy, vz = vx + ax, vy + ay, vz + az
    x, y, z = x + vx, y + vy, z + vz
    return (x, y, z), (vx, vy, vz), (ax, ay, az)

for i in range(10000):
    c = Counter(x[0] for x in particles)
    particles = [run_p(x) for x in particles if c[x[0]] == 1]
    print(len(particles))
