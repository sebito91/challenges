""" Module to find the largest product of a series of digits """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

def largest_product(series, run):
    """ function to find the largest product in a string of nums given """
    if run > len(series):
        raise ValueError("run is longer than series given")

    if run < 0:
        raise ValueError("run cannot be negative")

    if not series or not run:
        return 1

    if run == len(series):
        return reduce(lambda a, b: int(a)*int(b), list(series))

    perms = [series[i:(i + run)] for i in xrange(len(series) + 1 - run)]
    return max([reduce(lambda a, b: int(a)*int(b), list(series)) for series in perms])
