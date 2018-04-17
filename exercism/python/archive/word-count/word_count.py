""" Module that counts the number of words in a phrase """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals
from collections import defaultdict
import re

def word_count(src=''):
    """ Count the number of words in src, return dict with counts """
    src = src.lower().decode('utf-8')
    output = defaultdict(int)

    for each in [word for word in re.split(r'(?u)\W+|[_,]', src) if word]:
        output[each.lower()] += 1

    return output
