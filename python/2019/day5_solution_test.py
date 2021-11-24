#!/usr/bin/env python3

from unittest.mock import patch
import unittest
from io import StringIO
from day5_solution import IntcodeComputer

class TestDay5Solution(unittest.TestCase):

    @patch('builtins.input', lambda p: 77)
    def test_input_op(self):
        computer = IntcodeComputer([3,1,99])
        computer.run()
        self.assertEqual(computer.integers, [3,77,99])

    @patch('sys.stdout', new_callable=StringIO)
    def test_output_op(self, mock_stdout):
        computer = IntcodeComputer([4,2,99])
        computer.run()
        print(f"mock_stdout={mock_stdout}")

    def test_add(self):
        computer = IntcodeComputer([1,0,0,0,99])
        computer.run()
        self.assertEqual(computer.integers, [2,0,0,0,99])

    def test_multiply(self):
        computer = IntcodeComputer([2,3,0,3,99])
        computer.run()
        self.assertEqual(computer.integers, [2,3,0,6,99])

    def test_multiply2(self):
        computer = IntcodeComputer([2,4,4,5,99,0])
        computer.run()
        self.assertEqual(computer.integers, [2,4,4,5,99,9801])

    def test_multilple_ops(self):
        computer = IntcodeComputer([1,1,1,4,99,5,6,0,99])
        computer.run()
        self.assertEqual(computer.integers, [30,1,1,4,2,5,6,0,99])

    def test_multiple_mixed_parameter_modes(self):
        computer = IntcodeComputer([1002,4,3,4,33])
        computer.run()
        self.assertEqual(computer.integers, [1002,4,3,4,99])

if __name__ == '__main__':
    unittest.main()
