#!/usr/bin/env python3

from utils import read_input_raw
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
        if not x.append(int(n)):
            return n
    raise Exception("Couldn't find solution")

def part2(data):
    count = 0
    return count

if __name__ == '__main__':
    input_file = os.path.join(os.path.dirname(__file__), 'day9_input')
    data = read_input_raw(input_file)
    print(part1(data, 25))
    # print(part2(data))
