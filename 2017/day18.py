class Program(object):
    def __init__(self, program_id):
        self.id = program_id
        self.registers = {}
        self.registers['p'] = self.id
        self.receive_queue = []
        self.peer = None
        self.blocked = False
        self.count_send = 0
        self.current_instruction = 0

    def set_peer(self, peer):
        self.peer = peer

    def _get_val(self, x):
        try:
            return self.registers[x]
        except KeyError:
            try:
                return int(x)
            except ValueError:
                return 0

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

    def snd(self, x):
        """
        sends the value of x to its peer
        """
        self.count_send += 1
        self.peer.queue_message(self._get_val(x))

    def queue_message(self, val):
        self.receive_queue.append(val)
        self.blocked = False

    def rcv(self, x):
        """
        receives a value and stores it in register x.
        does nothing else while hasn't received anything.
        """
        if len(self.receive_queue) > 0:
            self.registers[x] = self.receive_queue.pop(0)
            self.blocked = False
        else:
            self.blocked = True
            return 0

    def finished(self, instructions):
        return (
            self.current_instruction < 0 or
            self.current_instruction >= len(instructions)
        )

    def jgz(self, x, y):
        """
        jumps with an offset of the value of Y, but only if the value of X is
        greater than zero. (An offset of 2 skips the next instruction, an
        offset of -1 jumps to the previous instruction, and so on.)
        """
        return self._get_val(y) if self._get_val(x) > 0 else None

    def run_instruction(self, instructions):
        instruction = instructions[self.current_instruction]
        method = getattr(self, instruction[0])
        offset = method(*instruction[1:])
        self.current_instruction += 1 if offset is None else offset


def run_programs(instructions):
    prog1 = Program(0)
    prog2 = Program(1)
    prog1.set_peer(prog2)
    prog2.set_peer(prog1)
    blocked = False
    current_program = prog1
    while not blocked:
        current_program.run_instruction(instructions)
        blocked = current_program.blocked or current_program.finished(instructions)
        if blocked:
            current_program = current_program.peer
            blocked = current_program.blocked or current_program.finished(instructions)
    print(prog2.count_send)


test_input = [
    ['snd', 1],
    ['snd', 2],
    ['snd', 'p'],
    ['rcv', 'a'],
    ['rcv', 'b'],
    ['rcv', 'c'],
    ['rcv', 'd']
]

print("run tests")
run_programs(test_input)


input_data = [
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
]
print("run prog")
run_programs(input_data)
