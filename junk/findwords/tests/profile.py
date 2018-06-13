#!/usr/bin/python

##rom guppy import hpy
import time

from ..wordtree import WordTree

start = time.time()

with open('dictionary.txt') as f:
    tree = WordTree(f.readlines())

end = time.time()

print (end - start)
