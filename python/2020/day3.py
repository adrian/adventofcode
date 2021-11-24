#!/usr/bin/env python3

from utils import read_input_raw
import sys
import os

def count_trees(map, x_step, y_step):
    x = y = trees = 0
    map_width = len(map[0]) - 1
    while y < len(map) - 1:
        x = (x + x_step) % map_width
        y = y + y_step
        if map[y][x] == '#':
            trees = trees + 1
    return trees

def part1_solution(map):
    return count_trees(map, 3, 1)

def part2_solution(map):
    a = count_trees(map, 1, 1)
    b = count_trees(map, 3, 1)
    c = count_trees(map, 5, 1)
    d = count_trees(map, 7, 1)
    e = count_trees(map, 1, 2)
    return a * b * c * d * e

if __name__ == '__main__':
    input_file = os.path.join(os.path.dirname(__file__), 'day3_input')
    # print(part1_solution(read_input_raw(input_file)))
    print(part2_solution(read_input_raw(input_file)))
