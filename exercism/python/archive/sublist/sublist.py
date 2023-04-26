"""Module to find a list within a list."""
# -*- coding: utf-8 -*-

from __future__ import unicode_literals


SUBLIST, SUPERLIST, EQUAL, UNEQUAL = 0, 1, 2, 3


def looper(one, two, len_one):
    """Check for completeness in the loops."""
    if any(all(a == b for (a, b) in zip(two[each:each + len_one], one)) for each in [i for i, x in enumerate(two) if x == one[0]]):
        return True

    return False


def sublist(one, two):
    """Check two lists."""
    len_one, len_two = len(one), len(two)

    if len_one == len_two and all(a == b for (a, b) in zip(one, two)):
        return EQUAL

    if len_one < len_two:
        if len_one == 0 or looper(one, two, len_one):
            return SUBLIST

    if len_one > len_two:
        if len_two == 0 or looper(two, one, len_two):
            return SUPERLIST

    return UNEQUAL
