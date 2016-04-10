#!/usr/bin/python

class WordTree(object):

    def __init__(self, entries):
        self.root = Root()
        self._entry_count = 0
        for entry in entries:
            self._add(entry)

    def entries(self):
        return self._entry_count

    def is_word(self, word):
        return self.root.is_word(word)

    def is_prefix(self, prefix):
        return self.root.is_prefix(prefix)

    def _add(self, entry):
        self._entry_count += 1
        self.root.insert(entry)

class Root(object):
    def __init__(self):
        self._children = []

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

    def _add_child(self, word):
        self._children.append(Node(word))

    def _find_child(self, value):
        for child in self._children:
            if child._value is value:
                return child
        return None

class Node(Root):
    def __init__(self, word):
        super(Node, self).__init__()
        self._value = word[0]
        if len(word) is 1:
            self._is_word = True
            print '{} true'.format(self._value)
        else:
            self._is_word = False
            print '{} {} false'.format(self._value, word[1:])
            self._add_child(word[1:])
        
    def value(self):
        return self._value
