#!/usr/bin/env python3

from utils import read_input_raw
import re
import sys

def count_occurances(letter, string):
    occurances = 0
    for c in string:
        if c == letter:
            occurances = occurances + 1
    return occurances

def part2_solution(password_data):
    correct_passwords = 0
    for pwd in password_data:
        parts = re.search("(\d+)-(\d+) (\w): (\w+)", pwd)
        pos1 = int(parts[1]) - 1
        pos2 = int(parts[2]) - 1
        letter = parts[3]
        password = parts[4]
        if (password[pos1] == letter and password[pos2] != letter) or (password[pos1] != letter and password[pos2] == letter):
            correct_passwords = correct_passwords + 1
    return correct_passwords

def part1_solution(password_data):
    correct_passwords = 0
    for pwd in password_data:
        parts = re.search("(\d+)-(\d+) (\w): (\w+)", pwd)
        min_len = int(parts[1])
        max_len = int(parts[2])
        letter = parts[3]
        password = parts[4]
        occurances = count_occurances(letter, password)
        if occurances >= min_len and occurances <= max_len:
            correct_passwords = correct_passwords + 1
    return correct_passwords

if __name__ == '__main__':
    #print(part1_solution(read_input_raw('day2_input')))
    print(part2_solution(read_input_raw('day2_input')))
