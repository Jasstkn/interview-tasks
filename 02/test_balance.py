import unittest
from balance import balance, calculate_balance

class TestStringMethods(unittest.TestCase):
    def test_balance(self):
        tests = [
            ([0, 0, 0], 0),
            ([1000, 200, 0], 1200),
            ([1000, 0, 1], 1010)
        ]
        for value, expected in tests:
            with self.subTest(value=value):
                self.assertEqual(balance(value[0], value[1], value[2]), expected)
    def test_calculate_balance(self):
        tests = [
            ([0, 0, 0, 0], 0),
            ([1000, 200, 0, 2], 1400),
            ([1000, 0, 5, 0], 1050),
            ([1000, 200, 2, 3], 1673.288)
        ]
        for value, expected in tests:
            with self.subTest(value=value):
                self.assertEqual(calculate_balance(value[0], value[1], value[2], value[3]), expected)

if __name__ == '__main__':
    unittest.main()
