#!/usr/bin/env python3

from utils import read_input_raw
import os
import re

class Instruction:
    def __init__(self, s):
        m = re.fullmatch(r"^(\w+) ([-+])(\d+)$", s)
        assert(m is not None)
        assert(len(m.groups()) == 3)
        self.instruction = m.groups()[0]
        if m.groups()[1] == '-':
            self.data = int(m.groups()[2]) * -1
        else:
            self.data = int(m.groups()[2])

class Computer:
    def __init__(self):
        self.accumulator = 0
        self.visited_instructions = []
        self.next_instruction_position = 0
    
    def reset(self):
        self.accumulator = 0
        self.visited_instructions = []
        self.next_instruction_position = 0

    def execute(self, instructions):
        while True:
            if self.next_instruction_position == len(instructions):
                return True # indicates the program terminated gracefully

            if self.next_instruction_position in self.visited_instructions:
                return False # indicates the program has entered an infinite loop

            self.visited_instructions.append(self.next_instruction_position)

            instruction = instructions[self.next_instruction_position]
            if instruction.instruction == "acc":
                self.accumulator = self.accumulator + instruction.data
                self.next_instruction_position = self.next_instruction_position + 1
            elif instruction.instruction == "nop":
                self.next_instruction_position = self.next_instruction_position + 1
            elif instruction.instruction == "jmp":
                self.next_instruction_position = self.next_instruction_position + instruction.data
            else:
                raise f"Invalid instruction {instruction.instruction}"

def parse_instructions(data):
    instructions = []
    for l in data:
        instructions.append(Instruction(l.rstrip()))
    return instructions

def part1(instructions):
    computer = Computer()
    computer.execute(instructions)
    return computer.accumulator

def part2(instructions):
    computer = Computer()

    for i in range(len(instructions)):
        if instructions[i].instruction == "jmp":
            instructions[i].instruction = "nop"
        elif instructions[i].instruction == "nop":
            instructions[i].instruction = "jmp"
        else:
            continue

        computer.reset()
        if computer.execute(instructions):
            return computer.accumulator

        if instructions[i].instruction == "nop":
            instructions[i].instruction = "jmp"
        else:
            instructions[i].instruction = "nop"

    raise Exception("Didn't find infinite loop instruction")

if __name__ == '__main__':
    input_file = os.path.join(os.path.dirname(__file__), 'day8_input')
    data = read_input_raw(input_file)
    instructions = parse_instructions(data)

    print(f"Part 1: {part1(instructions)}")

    print(f"Part 2: {part2(instructions)}")
