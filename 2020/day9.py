#!/usr/bin/env python3

from utils import read_input
import os

class XMAS:
    def __init__(self, num_prev_used_for_algorithom):
        self.recieved_numbers = []
        self.num_prev_used_for_algorithom = num_prev_used_for_algorithom

    def append(self, n):
        if len(self.recieved_numbers) < self.num_prev_used_for_algorithom:
            self.recieved_numbers.append(n)
            return True
        
        active_pool = self.recieved_numbers[len(self.recieved_numbers) - self.num_prev_used_for_algorithom:]
        for i in active_pool:
            if n - i in active_pool:
                self.recieved_numbers.append(n)
                return True

        return False

def part1(data, num_prev_used_for_algorithom):
    x = XMAS(num_prev_used_for_algorithom)
    for n in data:
        if not x.append(n):
            return n
    raise Exception("Couldn't find solution")

def find_valid_set(data, j, target):
    working_sum = data[j]
    for k in range(j + 1, len(data)):
        working_sum = working_sum + data[k]
        if working_sum > target:
            return []
        elif working_sum == target:
            return data[j:k + 1]
        else:
            continue
    raise Exception("valid_starting_position fell out of 'for' loop")

def part2(data, target):
    for i in range(0, len(data)):
        valid_set = find_valid_set(data, i, target)
        if valid_set:
            m1 = min(valid_set)
            m2 = max(valid_set)
            return m1 + m2
    raise Exception("Couldn't find valid set for part2")

if __name__ == '__main__':
    input_file = os.path.join(os.path.dirname(__file__), 'day9_input')
    data = read_input(input_file)
    part1_solution = part1(data, 25)
    # part1_solution = part1(data, 5)
    print(f"Part 1: {part1_solution}")
    print(f"Part 2: {part2(data, part1_solution)}")
