import unittest
import day7
import os
from utils import read_input_raw

class TestDay7(unittest.TestCase):

    def test_load_bags(self):
        input_file = os.path.join(os.path.dirname(__file__), 'day7_input_test')
        data = read_input_raw(input_file)
        bags = day7.load_bags(data)
        self.assertEqual(len(bags.keys()), 9)

        self.assertTrue('shiny gold' in bags)
        shiny_gold_bag = bags['shiny gold']
        self.assertEqual(len(shiny_gold_bag.edges), 2)

        for edge in shiny_gold_bag.edges:
            if edge.node.name == "dark olive":
                self.assertEqual(edge.weight, 1)
                self.assertEqual(edge.node, bags['dark olive'])
                self.assertEqual(len(edge.node.edges), 2)
            else:
                self.assertTrue(edge.node.name == "vibrant plum")
                self.assertEqual(edge.weight, 2)
                self.assertEqual(edge.node, bags['vibrant plum'])
                self.assertEqual(len(edge.node.edges), 2)

        self.assertTrue('faded blue' in bags)
        faded_blue_bag = bags['faded blue']
        self.assertEqual(len(faded_blue_bag.edges), 0)

    def _test_node(self):
        node1 = day7.Node("striped white")
        node2 = day7.Node("drab silver")
        node1.add_edge(node2, 4)

if __name__ == '__main__':
    unittest.main()
