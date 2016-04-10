#!/usr/bin/python

import unittest

from .. import wordtree

class TestWordTree(unittest.TestCase):

    def test_empty_dictionary(self):
        tree = wordtree.WordTree([])
        self.assertEqual(0, tree.entries())

    def test_count_one_entry(self):
        tree = wordtree.WordTree(['as'])
        self.assertEqual(1, tree.entries())

    def test_word_one_entry(self):
        tree = wordtree.WordTree(['as'])
        self.assertTrue(tree.is_word('as'))

    def test_not_word_prefix_one_entry(self):
        tree = wordtree.WordTree(['as'])
        self.assertFalse(tree.is_word('a'))

    def test_not_word_suffix_one_entry(self):
        tree = wordtree.WordTree(['as'])
        self.assertFalse(tree.is_word('astatine'))

    def test_prefix_one_entry(self):
        tree = wordtree.WordTree(['as'])
        self.assertTrue(tree.is_prefix('a'))

    def test_not_prefix_one_entry(self):
        tree = wordtree.WordTree(['as'])
        self.assertFalse(tree.is_prefix('z'))
