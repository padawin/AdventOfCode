class Duet(object):
    def _get_val(self, x):
        return self.registers[x] if x in self.registers else x

    def snd(self, x):
        """
        plays a sound with a frequency equal to the value of X.
        """
        self.last_sounded = x, self.registers[x]

    def set(self, x, y):
        """
        sets register X to the value of Y.
        """
        self.registers[x] = self._get_val(y)

    def add(self, x, y):
        """
        increases register X by the value of Y.
        """
        self.registers[x] += self._get_val(y)

    def mul(self, x, y):
        """
        sets register X to the result of multiplying the value contained in
        register X by the value of Y.
        """
        self.registers[x] *= self._get_val(y)

    def mod(self, x, y):
        """
        sets register X to the remainder of dividing the value contained in
        register X by the value of Y (that is, it sets X to the result of X
        modulo Y).
        """
        self.registers[x] = self.registers[x] % self._get_val(y)

    def rcv(self, x):
        """
        recovers the frequency of the last sound played, but only when the
        value of X is not zero. (If it is zero, the command does nothing.)
        """
        if self.registers[x] != 0:
            self.stop_process = True
        else:
            self.current_instruction += 1

    def jgz(self, x, y):
        """
        jumps with an offset of the value of Y, but only if the value of X is
        greater than zero. (An offset of 2 skips the next instruction, an
        offset of -1 jumps to the previous instruction, and so on.)
        """
        return self._get_val(y) if self.registers[x] > 0 else 1

    def run(self, duet):
        self.last_sounded = None
        self.stop_process = False
        self.registers = {
            'a': 0, 'b': 0, 'c': 0, 'd': 0, 'e': 0, 'f': 0, 'g': 0, 'h': 0, 'i': 0,
            'j': 0, 'k': 0, 'l': 0, 'm': 0, 'n': 0, 'o': 0, 'p': 0, 'q': 0, 'r': 0,
            's': 0, 't': 0, 'u': 0, 'v': 0, 'w': 0, 'x': 0, 'y': 0, 'z': 0
        }
        self.current_instruction = 0
        while not self.stop_process and self.current_instruction < len(duet):
            instruction = duet[self.current_instruction]
            method = getattr(self, instruction[0])
            next = method(*instruction[1:])
            self.current_instruction += next if next else 1


d = Duet()
d.run([
    ['set', 'a', 1],
    ['add', 'a', 2],
    ['mul', 'a', 'a'],
    ['mod', 'a', 5],
    ['snd', 'a'],
    ['set', 'a', 0],
    ['rcv', 'a'],
    ['jgz', 'a', -1],
    ['set', 'a', 1],
    ['jgz', 'a', -2]
])
assert d.last_sounded == ('a', 4)


d.run([
    ['set', 'i', 31],
    ['set', 'a', 1],
    ['mul', 'p', 17],
    ['jgz', 'p', 'p'],
    ['mul', 'a', 2],
    ['add', 'i', -1],
    ['jgz', 'i', -2],
    ['add', 'a', -1],
    ['set', 'i', 127],
    ['set', 'p', 622],
    ['mul', 'p', 8505],
    ['mod', 'p', 'a'],
    ['mul', 'p', 129749],
    ['add', 'p', 12345],
    ['mod', 'p', 'a'],
    ['set', 'b', 'p'],
    ['mod', 'b', 10000],
    ['snd', 'b'],
    ['add', 'i', -1],
    ['jgz', 'i', -9],
    ['jgz', 'a', 3],
    ['rcv', 'b'],
    ['jgz', 'b', -1],
    ['set', 'f', 0],
    ['set', 'i', 126],
    ['rcv', 'a'],
    ['rcv', 'b'],
    ['set', 'p', 'a'],
    ['mul', 'p', -1],
    ['add', 'p', 'b'],
    ['jgz', 'p', 4],
    ['snd', 'a'],
    ['set', 'a', 'b'],
    ['jgz', 1, 3],
    ['snd', 'b'],
    ['set', 'f', 1],
    ['add', 'i', -1],
    ['jgz', 'i', -11],
    ['snd', 'a'],
    ['jgz', 'f', -16],
    ['jgz', 'a', -19],
])
print(d.last_sounded)
