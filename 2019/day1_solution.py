#!/usr/bin/env python

from functools import reduce
from math import floor

with open("day1_input") as f:
  modules_mass = [int(line) for line in f.readlines()]

def fuel_required(module_mass):
  return int(floor(module_mass / 3)) - 2

fuel_per_module = [fuel_required(mass) for mass in modules_mass]

print sum(fuel_per_module)
