""" Module to return the set of prime factors for a number """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

import math

def prime_factors(number):
    """ return the list of prime factors of the given number """
    if number == 1:
        return []

    thelist = []
    while number % 2 == 0:
        number /= 2
        thelist.append(2)

    for each in xrange(3, int(math.sqrt(number))+1, 2):
        while number % each == 0:
            number /= each
            thelist.append(each)

    if number > 2:
        thelist.append(number)

    return thelist
