#!/usr/bin/env python3

import unittest

from day2_solution import IntcodeComputer

class TestDay2Solution(unittest.TestCase):

    def test1(self):
        computer = IntcodeComputer([1,0,0,0,99])
        computer.run()
        self.assertEqual(computer.integers, [2,0,0,0,99])

    def test2(self):
        computer = IntcodeComputer([2,3,0,3,99])
        computer.run()
        self.assertEqual(computer.integers, [2,3,0,6,99])

    def test3(self):
        computer = IntcodeComputer([2,4,4,5,99,0])
        computer.run()
        self.assertEqual(computer.integers, [2,4,4,5,99,9801])

    def test4(self):
        computer = IntcodeComputer([1,1,1,4,99,5,6,0,99])
        computer.run()
        self.assertEqual(computer.integers, [30,1,1,4,2,5,6,0,99])

if __name__ == '__main__':
    unittest.main()
