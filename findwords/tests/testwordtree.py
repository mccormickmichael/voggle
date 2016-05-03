#!/usr/bin/python

import unittest

from .. import wordtree


class TestWordTree(unittest.TestCase):

    def test_empty_dictionary(self):
        tree = wordtree.WordTree([])
        self.assertEqual(0, tree.count())

    def test_count_one_entry(self):
        tree = wordtree.WordTree(['as'])
        self.assertEqual(1, tree.count())

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

    def test_count_two_entries(self):
        tree = wordtree.WordTree(['as', 'at'])
        self.assertEquals(2, tree.count())

    def test_word_two_entries(self):
        tree = wordtree.WordTree(['as', 'at'])
        self.assertTrue(tree.is_word('as'))
        self.assertTrue(tree.is_word('at'))

    def test_longer_words(self):
        tree = wordtree.WordTree(['astatine', 'asterisk', 'at', 'as', 'ass'])
        self.assertTrue(tree.is_word('as'))
        self.assertTrue(tree.is_word('at'))
        self.assertTrue(tree.is_word('ass'))
        self.assertFalse(tree.is_word('astatin'))

    def test_longer_prefixes(self):
        tree = wordtree.WordTree(['astatine', 'asterisk', 'at', 'as', 'ass'])
        self.assertTrue(tree.is_prefix('as'))
        self.assertTrue(tree.is_prefix('astatin'))
        self.assertFalse(tree.is_prefix('astatinf'))
        self.assertFalse(tree.is_prefix('astatines'))
