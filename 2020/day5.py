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
        col = num_from_binary_chars(seat_chars[7:10], 'R')
        seats.append((row, col))
    return seats

def seat_ids(seats):
    return [(seat[0] * 8 + seat[1]) for seat in seats]

def missing_seat_ids(seat_ids):
    m = []
    for row in range(0, 128):
        for col in range(0, 8):
            seat_id = row * 8 + col
            if seat_id not in seat_ids:
                m.append(seat_id)
                print(f"Row: {row}, Col: {col}, Seat ID: {seat_id}")
    return m

if __name__ == '__main__':
    input_file = os.path.join(os.path.dirname(__file__), 'day5_input')
    seats = load_seats(read_input_raw(input_file))
    # print(max(seat_ids(seats)))

    m = missing_seat_ids(seat_ids(seats))
    # picked the odd looking missing seat in the output, 659
