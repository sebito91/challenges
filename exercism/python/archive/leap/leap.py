""" module to calculate whether this is a leap year """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

def is_leap_year(year=0):
    """ method to check for the leap year """
    if year == 0:
        return False

    if year % 4 == 0 and (year % 100 != 0 or year % 400 == 0):
        return True

    return False
