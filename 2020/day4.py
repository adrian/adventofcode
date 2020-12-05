#!/usr/bin/env python3

from utils import read_input_raw
import sys
import os
import logging
import re

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
    for field in ['byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid']:
        if field not in passport:
            return False
    return True

def validate_num(s, expected_len, min_val, max_val):
    if len(s) != expected_len:
        return False
    num = int(s)
    return num >= min_val and num <= max_val

def part2_valid(passport):
    if not part1_valid(passport):
        print(f"missing required fields: {passport}")
        return False

    if not validate_num(passport['byr'], 4, 1920, 2002):
        print(f"invalid byr: {passport['byr']}")
        return False

    if not validate_num(passport['iyr'], 4, 2010, 2020):
        print(f"invalid iyr: {passport['iyr']}")
        return False

    if not validate_num(passport['eyr'], 4, 2020, 2030):
        print(f"invalid eyr: {passport['eyr']}")
        return False

    hgt = passport['hgt']
    m = re.fullmatch(r"(\d+)(cm|in)", hgt)
    if m == None or len(m.groups()) != 2:
        print(f"invalid hgt: {hgt}")
        return False
    hgt_num = int(m.group(1))
    if m.group(2) == "cm":
        if hgt_num < 150 or hgt_num > 193:
            print(f"invalid hgt: {hgt}")
            return False
    else:
        if hgt_num < 59 or hgt_num > 76:
            print(f"invalid hgt: {hgt}")
            return False

    hcl = passport['hcl']
    if not re.fullmatch(r"#[0123456789abcdef]{6}", hcl):
        print(f"invalid hcl: {hcl}")
        return False

    ecl = passport['ecl']
    if not ecl in ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth']:
        print(f"invalid ecl: {ecl}")
        return False

    pid = passport['pid']
    if not re.fullmatch(r"[0123456789]{9}", pid):
        print(f"invalid pid: {pid}")
        return False

    return True

if __name__ == '__main__':
    input_file = os.path.join(os.path.dirname(__file__), 'day4_input')
    passports = load_passports(read_input_raw(input_file))
    # valid_passports = [passport for passport in passports if part1_valid(passport)]
    valid_passports = [passport for passport in passports if part2_valid(passport)]
    print(len(valid_passports))
