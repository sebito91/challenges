""" Module to find a list within a list """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

SUBLIST, SUPERLIST, EQUAL, UNEQUAL = 0, 1, 2, 3

def check_lists(one, two):
    """ Function to check two lists """
    len_one, len_two = len(one), len(two)

    if len_one == len_two and all([a == b for (a, b) in zip(one, two)]):
        return EQUAL
#
#    if len_one < len_two and any([a != b for (a, b) in zip(two[each:each + len_one], one) \
#            for each in [i for i, x in enumerate(two) if x == one[0]]]):
#        return SUBLIST

    print [i for i, x in enumerate(two) if x == one[0]]
    print [(a, b) for each in [i for i, x in enumerate(two) if x == one[0]] \
            for (a, b) in zip(two[each:each + len_one], one)]
    if len_one < len_two and all([a == b for each in [i for i, x in enumerate(two) if x == one[0]] \
            for (a, b) in zip(one, two[each:each + len_one])]):
        return SUBLIST

    if len_one > len_two and all([a == b for each in [i for i, x in enumerate(one) if x == two[0]] \
            for (a, b) in zip(two, one[each:each + len_two])]):
        return SUPERLIST

    return UNEQUAL
