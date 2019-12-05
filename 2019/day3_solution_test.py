#!/usr/bin/env python3

import unittest

from day3_solution import find_intersections, shortest_manhattan_distance, \
    manhattan_distance

class TestDay3Solution(unittest.TestCase):

    def test_shortest_manhattan_length1(self):
        route1 = "R8,U5,L5,D3"
        route2 = "U7,R6,D4,L4"

        self.assertEqual(
            6,
            manhattan_distance(
                shortest_manhattan_distance(
                    find_intersections(route1, route2)
            ))
        )


    def test_shortest_manhattan_length2(self):
        route1 = "R75,D30,R83,U83,L12,D49,R71,U7,L72"
        route2 = "U62,R66,U55,R34,D71,R55,D58,R83"

        self.assertEqual(
            159,
            manhattan_distance(
                shortest_manhattan_distance(
                    find_intersections(route1, route2)
            ))
        )

    def test_shortest_manhattan_length3(self):
        route1 = "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"
        route2 = "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"

        self.assertEqual(
            135,
            manhattan_distance(
                shortest_manhattan_distance(
                    find_intersections(route1, route2)
            ))
        )

if __name__ == '__main__':
    unittest.main()
