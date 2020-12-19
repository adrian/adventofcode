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
            inner_bag =  bags[inner_bag_name] if inner_bag_name in bags.keys() else Node(inner_bag_name)
            if inner_bag_name not in bags.keys():
                inner_bag = Node(inner_bag_name)
                bags[inner_bag_name] = inner_bag
            # inner_bag.add_edge(outer_bag, int(bag_desc[0]))
            outer_bag.add_edge(inner_bag, int(bag_desc[0]))
    return bags

def print_bag(bag, depth):
    spaces = " " * depth
    print(f"{spaces}{bag.name}:")
    for bag_edge in bag.edges:
        outer_bag = bag_edge.node
        print_bag(outer_bag, depth + 1)

def explore_parent_bags(bag, visited_bags):
    for bag_edge in bag.edges:
        outer_bag = bag_edge.node
        if outer_bag.name not in visited_bags:
            visited_bags.append(outer_bag.name)
            explore_parent_bags(outer_bag, visited_bags)

def number_of_bags_in(bag):
    total = 0
    for bag_edge in bag.edges:
        total = total + bag_edge.weight + (bag_edge.weight * number_of_bags_in(bag_edge.node))
    return total

def part1(data):
    bags = load_bags(data)

    shiny_gold_bag = bags['shiny gold']
    # print_bag(shiny_gold_bag, 0)

    visited_bags = []
    explore_parent_bags(shiny_gold_bag, visited_bags)

    return len(visited_bags)

def part2(data):
    bags = load_bags(data)
    shiny_gold_bag = bags['shiny gold']
    return number_of_bags_in(shiny_gold_bag)

if __name__ == '__main__':
    input_file = os.path.join(os.path.dirname(__file__), 'day7_input')
    data = read_input_raw(input_file)

    # My "map" is uni-directional. To get part 1 working I need to swap
    # the direction by switching the commented out line in load_bags().
    # print(part1(data))

    print(part2(data))
