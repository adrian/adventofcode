#!/usr/bin/env python3

import re
from utils import read_input_raw

class Point:
    def __init__(self, x=0, y=0):
        self.x = x
        self.y = y

    def __lt__(self, other):
        return manhattan_distance(self) < manhattan_distance(other)

    def __eq__(self, other):
        return self.x == other.x and self.y == other.y

    def __str__(self):
        return f"(x={self.x},y={self.y})"

    __repr__ = __str__

    def __hash__(self):
        return hash((self.x, self.y))

class Instruction:
    def __init__(self, s):
        search_results = re.search('([RLUD])(\d+)', s)
        self.direction = search_results.group(1)
        self.steps = int(search_results.group(2))

    def __str__(self):
        return f"{self.direction}{self.steps}"

def step(point, direction):
    if direction == 'R':
        return Point(point.x + 1, point.y)
    elif direction == 'L':
        return Point(point.x - 1, point.y)
    elif direction == 'U':
        return Point(point.x, point.y + 1)
    elif direction == 'D':
        return Point(point.x, point.y - 1)
    else:
        raise Exception(f"Unknown direction: {direction}")

def find_intersections(route1_s, route2_s):
    route1 = [Instruction(inst) for inst in route1_s.split(',')]
    route2 = [Instruction(inst) for inst in route2_s.split(',')]

    intersections = []

    r1_current_point = Point()
    route1_points_visited = set()
    for r1_instruction in route1:
        for _ in range(r1_instruction.steps, 0, -1):
            r1_current_point = step(r1_current_point, r1_instruction.direction)
            route1_points_visited.add(r1_current_point)

    r2_current_point = Point()
    for r2_instruction in route2:
        for _ in range(r2_instruction.steps, 0, -1):
            r2_current_point = step(r2_current_point, r2_instruction.direction)
            if r2_current_point in route1_points_visited:
                intersections.append(r2_current_point)

    return intersections

def manhattan_distance(point):
    return abs(point.x) + abs(point.y)

def shortest_manhattan_distance(points):
    return min(points, key=lambda p: abs(p.x) + abs(p.y))

if __name__ == '__main__':
    lines = read_input_raw('day3_input')
    route1 = lines[0]
    route2 = lines[1]
    d = manhattan_distance(
            shortest_manhattan_distance(
                find_intersections(route1, route2)))
    print(f"Shortest Manhattan Distance: {d}")
