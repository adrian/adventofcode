#!/usr/bin/env python3

import unittest

from day4_solution import passwords_in_range, has_twins, digits_increasing

class TestDay4Solution(unittest.TestCase):

    def test1(self):
        self.assertEqual(1, passwords_in_range(111111, 111111))
        self.assertEqual(0, passwords_in_range(223450, 223450))
        self.assertEqual(0, passwords_in_range(123789, 123789))
        self.assertEqual(1, passwords_in_range(122345, 122345))

    def test_has_twins(self):
        self.assertEqual(True, has_twins('111111'))
        self.assertEqual(False, has_twins('1234567'))
        self.assertEqual(True, has_twins('112345'))
        self.assertEqual(True, has_twins('123455'))
        self.assertEqual(True, has_twins('1233456'))

    def test_digits_increasing(self):
        self.assertEqual(True, digits_increasing('123456'))
        self.assertEqual(True, digits_increasing('11234'))
        self.assertEqual(False, digits_increasing('123454'))

if __name__ == '__main__':
    unittest.main()
