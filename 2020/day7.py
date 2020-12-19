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
        print(f"Processing {outer_bag_name}")
        outer_bag = bags[outer_bag_name] if outer_bag_name in bags.keys() else Node(outer_bag_name)
        if outer_bag_name not in bags.keys():
            bags[outer_bag_name] = outer_bag
        bag_desc_parts = re.findall(r"(\d+) ([\w ]+) bags?", parts[1])
        for bag_desc in bag_desc_parts:
            inner_bag_name = bag_desc[1]
            inner_bag = None
            if inner_bag_name in bags.keys():
                print(f"Reusing {inner_bag_name} Node")
                inner_bag = bags[inner_bag_name]
            else:
                print(f"Creating a new Node for {inner_bag_name}")
                inner_bag = Node(inner_bag_name)
                bags[inner_bag_name] = inner_bag
            print(f"Added edge from {outer_bag_name} to {inner_bag.name}")
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
        print(f"{bag.name} is within {outer_bag.name}")
        if outer_bag.name not in visited_bags:
            visited_bags.append(outer_bag.name)
            explore_parent_bags(outer_bag, visited_bags)

def part1(data):
    bags = load_bags(data)

    print(bags)
    print(f"There are {len(bags.keys())} bag colors")

    shiny_gold_bag = bags['shiny gold']
    print_bag(shiny_gold_bag, 0)

    # I think I've got the node pointer backwards
    # shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
    # The pointers should go from dark olive to shiny gold
    # I do them the other way round

    visited_bags = []
    explore_parent_bags(shiny_gold_bag, visited_bags)

    visited_bags.sort()
    print(visited_bags)
    return len(visited_bags)

def part2(data):
    count = 0
    return count

if __name__ == '__main__':
    input_file = os.path.join(os.path.dirname(__file__), 'day7_input_test')
    data = read_input_raw(input_file)
    print(part1(data))
    # print(part2(data))
