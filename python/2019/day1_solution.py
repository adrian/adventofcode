#!/usr/bin/env python3

from math import floor
from utils import read_input

def fuel_required(mass):
    return int(floor(mass / 3)) - 2

def solution(masses):
    return sum([fuel_required(mass) for mass in masses])

if __name__ == '__main__':
    print(solution(read_input('day1_input')))
