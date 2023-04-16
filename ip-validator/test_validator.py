import unittest
from validator import ip_validator

class TestStringMethods(unittest.TestCase):
    def test_ip_validator(self):
        tests = [
            ('0.0.0.1', True),
            ('192.168.0.1', True),
            ('192.168.0.A', False),
            ('', False),
            ('256.0.0.1', False),
            ('001.0.0.1', False),
            ('001.0.0', False),
        ]
        for value, expected in tests:
            with self.subTest(value=value):
                self.assertEqual(ip_validator(value), expected)

if __name__ == '__main__':
    unittest.main()
