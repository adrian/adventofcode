#!/usr/bin/env python3

from utils import read_input

def part1_solution(numbers):
    for i, num1 in enumerate(numbers):
        for j, num2 in enumerate(numbers, start=i + 1):
            if num1 + num2 == 2020:
                return num1 * num2

def part2_solution(numbers):
    for i, num1 in enumerate(numbers):
        for j, num2 in enumerate(numbers, start=i + 1):
            for k, num3 in enumerate(numbers, start=j + 1):
                if num1 + num2 + num3 == 2020:
                    return num1 * num2 * num3

if __name__ == '__main__':
    # print(part1_solution(read_input('day1_input')))
    print(part2_solution(read_input('day1_input')))
