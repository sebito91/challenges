""" Module to return the slices of numbers based on the given input """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

def slices(inputs, num):
    """ Function to retun the slices of consecutive numbers given inputs """
    if not inputs or not num:
        raise ValueError("no inputs")

    if len(inputs) < num:
        raise ValueError("len of inputs less than num requested")

    if len(inputs) == num:
        return [[int(i) for i in list(inputs)]]

    thelist = [list(inputs[each:each+num]) for each in range(len(inputs) - num + 1)]

    return [[int(i) for i in each] for each in thelist]
