#!/usr/bin/env python3

from math import floor
from utils import read_file_as_int_array

class IntcodeComputer:

    def __init__(self, integers):
        self.integers = integers

    def run(self):
        position = 0
        while True:
            # print(f"Instruction: {self.integers[position]}," \
            #       f"{self.integers[position + 1]}," \
            #       f"{self.integers[position + 2]}," \
            #       f"{self.integers[position + 3]}")

            if self.integers[position] > 99:
                # ASSUMPTION: one 2 digit instruction and 3 parameters
                s = str(self.integers[position]).zfill(5)
                parameter3_mode = int(s[0:1])
                parameter2_mode = int(s[1:2])
                parameter1_mode = int(s[2:3])
                instruction = int(s[3:5])
            else:
                instruction = self.integers[position]
                parameter1_mode = 0
                parameter2_mode = 0
                parameter3_mode = 0

            if instruction == 99:
                break
            elif instruction in [1, 2]:
                num1_pos = self.integers[position + 1] if parameter1_mode == 0 \
                    else position + 1
                num2_pos = self.integers[position + 2] if parameter2_mode == 0 \
                    else position + 2
                result_pos = self.integers[position + 3]
                num1 = self.integers[num1_pos]
                num2 = self.integers[num2_pos]
                if instruction == 1:
                    self.integers[result_pos] = num1 + num2
                else:
                    self.integers[result_pos] = num1 * num2
                position = position + 4
            elif instruction == 3:
                save_to_pos = self.integers[position + 1]
                data = input("Enter Input: ")
                self.integers[save_to_pos] = int(data)
                position = position + 2
            elif instruction == 4:
                pos_to_print = self.integers[position + 1]
                print(f"Output: {self.integers[pos_to_print]}")
                position = position + 2
            else:
                raise Exception(f"Unknown instruction {instruction}")

if __name__ == '__main__':
    computer = IntcodeComputer(read_file_as_int_array('day5_input'))
    computer.run()
