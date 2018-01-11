""" Module to perform permutations of the difference
of squares of the first N natural numbers """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

import math


def square_of_sum(num):
    """ return the square of the sum of the first N natural numbers """
    return math.pow(sum([a for a in range(1, num+1) if a]), 2)


def sum_of_squares(num):
    """ return the sum of the squares of the first N natural numbers """
    return sum([math.pow(a, 2) for a in range(1, num+1) if a])


def difference(num):
    """ return the difference between the square_of_sum and sum_of_squares
    for the first N natural numbers
    """
    return square_of_sum(num) - sum_of_squares(num)
