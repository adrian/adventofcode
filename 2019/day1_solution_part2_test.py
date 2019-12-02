#!/usr/bin/env python3

import unittest

from day1_solution_part2 import solution

class TestDay1SolutionPart2(unittest.TestCase):

    def test1(self):
        self.assertEqual(solution([14]), 2)

    def test2(self):
        self.assertEqual(solution([1969]), 966)

    def test3(self):
        self.assertEqual(solution([100756]), 50346)

if __name__ == '__main__':
    unittest.main()
