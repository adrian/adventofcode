#!/usr/bin/env python3

from utils import read_input_raw
import sys
import os
import logging

def assemble_passport(lines):
    passport = {}
    for line in lines:
        fields = line.split(" ")
        for field in fields:
            keyval = field.split(":")
            assert(keyval[0] not in passport.keys())
            passport[keyval[0]] = keyval[1]
    return passport

def load_passports(raw_passport_data):
    passports = []
    single_passport_lines = []
    for line in raw_passport_data:
        line = line.rstrip()
        if line == "":
            passports.append(assemble_passport(single_passport_lines))
            single_passport_lines = []
        else:
            single_passport_lines.append(line)
    passports.append(assemble_passport(single_passport_lines))
    return passports

def part1_valid(passport):
    required_fields = ['byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid']
    for field in required_fields:
        if field not in passport:
            return False
    return True

def part1_solution(raw_passport_data):
    passports = load_passports(raw_passport_data)
    valid_passports = [passport for passport in passports if part1_valid(passport)]
    return len(valid_passports)

if __name__ == '__main__':
    input_file = os.path.join(os.path.dirname(__file__), 'day4_input')
    print(part1_solution(read_input_raw(input_file)))
