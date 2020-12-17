#!/usr/bin/env python3

from utils import read_input_raw
import os
import re

class Node:
    def __init__(self, name):
        self.name = name
        self.edges = []
    
    def add_edge(self, target_node, weight):
        self.edges.append(Edge(weight, target_node))

class Edge:
    def __init__(self, weight, node):
        self.weight = weight
        self.node = node

def load_bags(data):
    bags = {}
    for line in data:
        parts = line.split('bags contain')
        outer_bag_name = parts[0].rstrip()
        outer_bag = bags[outer_bag_name] if outer_bag_name in bags.keys() else Node(outer_bag_name)
        if outer_bag_name not in bags.keys():
            bags[outer_bag_name] = outer_bag
        bag_desc_parts = re.findall(r"(\d+) ([\w ]+) bags?", parts[1])
        for bag_desc in bag_desc_parts:
            inner_bag_name = bag_desc[1]
            inner_bag = bags[inner_bag_name] if inner_bag_name in bags.keys() else Node(inner_bag_name)
            if inner_bag_name not in bags.keys():
                bags[inner_bag_name] = inner_bag
            outer_bag.add_edge(inner_bag, int(bag_desc[0]))
    return bags

def part1(data):
    count = 0
    return count

def part2(data):
    count = 0
    return count

if __name__ == '__main__':
    input_file = os.path.join(os.path.dirname(__file__), 'day6_input')
    data = read_input_raw(input_file)
    print(part1(data))
    # print(part2(data))
