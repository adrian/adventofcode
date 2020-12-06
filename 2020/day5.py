#!/usr/bin/env python3

from utils import read_input_raw
import os
import math

def num_from_binary_chars(chars, high_char):
    num = 0
    l = len(chars)
    for i in range(0, l):
        if chars[i] == high_char:
            p = (l - 1) - i
            num = num + int(math.pow(2, p))
    return num

def load_seats(all_seat_chars):
    seats = []
    for seat_chars in all_seat_chars:
        row = num_from_binary_chars(seat_chars[:7], 'B')
        col = num_from_binary_chars(seat_chars[7:], 'B')
        seats.append((row, col))
    return seats

def seats_ids(seats):
    return [seat[0] * 8 + seat[1] for seat in seats]

if __name__ == '__main__':
    input_file = os.path.join(os.path.dirname(__file__), 'day5_input')
    seats = load_seats(read_input_raw(input_file))
    print(max(seats_ids(seats)))
