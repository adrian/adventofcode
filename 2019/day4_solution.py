#!/usr/bin/env python3

def has_twins(s):
    previous = None
    for t in s:
        if t == previous:
            return True
        previous = t
    return False

def digits_increasing(s):
    previous = 0
    for t in s:
        e = int(t)
        if e < previous:
            return False
        previous = e
    return True

def passwords_in_range(start, finish):
    pw_count = 0
    for i in range(start, finish + 1):
        s = str(i)
        if has_twins(s) and digits_increasing(s):
            pw_count = pw_count + 1
    return pw_count

if __name__ == '__main__':
    print(f"Number of passwords: {passwords_in_range(134792, 675810)}")
