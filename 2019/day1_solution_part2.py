#!/usr/bin/env python3

from math import floor
from utils import read_input

def fuel_required(starting_mass):
    mass = starting_mass
    total_fuel = 0
    while True:
        fuel = max(0, int(floor(mass / 3)) - 2)
        if fuel == 0:
            break
        total_fuel = total_fuel + fuel
        mass = fuel
    return total_fuel

def solution(masses):
    return sum([fuel_required(mass) for mass in masses])

if __name__ == '__main__':
    print(solution(read_input('day1_input')))
