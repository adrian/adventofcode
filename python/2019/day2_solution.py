#!/usr/bin/env python3

from math import floor
from utils import read_file_as_int_array

class IntcodeComputer:

    def __init__(self, integers):
        self.integers = integers

    def run(self):
        position = 0
        while True:
            instruction = self.integers[position]
            if instruction == 99:
                break
            elif instruction in [1, 2]:
                num1_pos = self.integers[position + 1]
                num2_pos = self.integers[position + 2]
                result_pos = self.integers[position + 3]
                num1 = self.integers[num1_pos]
                num2 = self.integers[num2_pos]
                if instruction == 1:
                    self.integers[result_pos] = num1 + num2
                else:
                    self.integers[result_pos] = num1 * num2
            elif instruction == 2:
                num1 = self.integers[position + 1]
                num2 = self.integers[position + 2]
                self.integers[position + 3] = num1 * num2
            else:
                raise Exception(f"Unknown instruction {instruction}")
            position = position + 4

if __name__ == '__main__':
    computer = IntcodeComputer(read_file_as_int_array('day2_input'))
    computer.integers[1] = 12
    computer.integers[2] = 2
    computer.run()
    print(computer.integers)
