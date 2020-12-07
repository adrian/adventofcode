#!/usr/bin/env python3

from utils import read_input_raw
import os

def group_answers_generator(data):
    group_lines = []
    for line in data:
        line = line.rstrip()
        if line == "":
            yield group_lines
            group_lines = []
        else:
            group_lines.append(line)
    yield group_lines

def num_unique_yeses(group_answers):
    yeses = {}
    all_answers = ''.join(group_answers)
    for answer in all_answers:
        yeses[answer] = True
    return len(yeses.keys())

def part1(data):
    count = 0
    for group_answers in group_answers_generator(data):
        count = count + num_unique_yeses(group_answers)
    return count

if __name__ == '__main__':
    input_file = os.path.join(os.path.dirname(__file__), 'day6_input')
    data = read_input_raw(input_file)
    print(part1(data))
