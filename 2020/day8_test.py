import unittest
import day8
import os

class TestDay8(unittest.TestCase):

    def test_simple_instruction(self):
        instruction = day8.Instruction("nop +0")
        self.assertEqual(instruction.instruction, "nop")
        self.assertEqual(instruction.data, 0)

    def test_instruction_with_negative_data(self):
        instruction = day8.Instruction("jmp -3")
        self.assertEqual(instruction.instruction, "jmp")
        self.assertEqual(instruction.data, -3)

    def test_instruction_with_large_number(self):
        instruction = day8.Instruction("acc +999")
        self.assertEqual(instruction.instruction, "acc")
        self.assertEqual(instruction.data, 999)


if __name__ == '__main__':
    unittest.main()
