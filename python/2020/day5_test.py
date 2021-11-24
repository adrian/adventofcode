import unittest
import day5

class TestDay5(unittest.TestCase):

    def test_row_from_bsp(self):
        self.assertEqual(day5.num_from_binary_chars('FBFBBFF', 'B'), 44)
        self.assertEqual(day5.num_from_binary_chars('BFFFBBF', 'B'), 70)
        self.assertEqual(day5.num_from_binary_chars('FFFBBBF', 'B'), 14)
        self.assertEqual(day5.num_from_binary_chars('BBFFBBF', 'B'), 102)

    def test_col_from_bsp(self):
        self.assertEqual(day5.num_from_binary_chars('RLR', 'R'), 5)
        self.assertEqual(day5.num_from_binary_chars('RRR', 'R'), 7)
        self.assertEqual(day5.num_from_binary_chars('LLL', 'R'), 0)

if __name__ == '__main__':
    unittest.main()
