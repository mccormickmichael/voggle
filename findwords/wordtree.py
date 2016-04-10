#!/usr/bin/python

class Node(object):
    def __init__(self, word):
        self._value = word[0]
        self._children = None
        if len(word) is 1:
            self._is_word = True
            print '{} true'.format(self._value)
        else:
            self._is_word = False
            print '{} {} false'.format(self._value, word[1:])
            self._add_child(word[1:])

    def insert(self, word):
        child_node = self._find_child(word[0])
        if child_node is None:
            self._add_child(word)
        elif len(word) is 1:
            child_node._is_word = True
        else:
            child_node.insert(word[1:])

    def is_prefix(self, prefix):
        child_node = self._find_child(prefix[0])
        if child_node is None:
            return False
        if len(prefix) is 1:
            return True
        return child_node.is_prefix(prefix[1:])

    def is_word(self, word):
        child_node = self._find_child(word[0])
        if child_node is None:
            return False
        if child_node._is_word and len(word) is 1:
            return True
        if len(word) is 1:
            return False
        return child_node.is_word(word[1:])

    def _ensure_children(self):
        if self._children is None:
            self._children = []

    def _add_child(self, word):
        self._ensure_children()
        self._children.append(Node(word))

    def _find_child(self, value):
        self._ensure_children()
        for child in self._children:
            if child._value is value:
                return child
        return None

class WordTree(Node):
    def __init__(self, entries):
        self._children = []
        self._entry_count = 0
        for entry in entries:
            self.insert(entry)

    def count(self):
        return self._entry_count

    def insert(self, entry):
        self._entry_count += 1
        super(WordTree, self).insert(entry)
