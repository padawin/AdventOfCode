import sys


class Program(object):
    def __init__(self):
        self.registers = {}
        self.count_mul = 0
        self.current_instruction = 0

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

    def sub(self, x, y):
        """
        sets register X to the result of multiplying the value contained in
        register X by the value of Y.
        """
        self.registers[x] = self._get_val(x) - self._get_val(y)

    def mul(self, x, y):
        """
        sets register X to the result of multiplying the value contained in
        register X by the value of Y.
        """
        self.count_mul += 1
        self.registers[x] = self._get_val(x) * self._get_val(y)

    def jnz(self, x, y):
        """
        jumps with an offset of the value of Y, but only if the value of X is
        greater than zero. (An offset of 2 skips the next instruction, an
        offset of -1 jumps to the previous instruction, and so on.)
        """
        return self._get_val(y) if self._get_val(x) != 0 else None

    def run_instruction(self, instructions):
        instruction = instructions[self.current_instruction].split(' ')
        method = getattr(self, instruction[0])
        offset = method(*instruction[1:])
        self.current_instruction += 1 if offset is None else offset


def run_program(instructions):
    instructions = [l.rstrip('\n') for l in instructions.readlines()]
    prog = Program()
    while 0 <= prog.current_instruction < len(instructions):
        prog.run_instruction(instructions)
    print(prog.count_mul)


run_program(sys.stdin)
