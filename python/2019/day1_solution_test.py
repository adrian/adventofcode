#!/usr/bin/env python3

import unittest

from day1_solution import solution

class TestDay1Solution(unittest.TestCase):

    def test1(self):
        self.assertEqual(solution([12]), 2)

    def test2(self):
        self.assertEqual(solution([14]), 2)

    def test3(self):
        self.assertEqual(solution([1969]), 654)

    def test4(self):
        self.assertEqual(solution([100756]), 33583)

if __name__ == '__main__':
    unittest.main()
