""" calculate the number of mutations in the DNA strand (hamming distance) """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

def distance(one='', two=''):
    """ calculate the difference between the two inputs """
    return len(''.join([a for a, b in zip(one, two) if a != b]))
