""" Module to return the number of grains """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

def on_square(num):
    """ return the number of grains on a given square """
    return long(2 ** (num - 1))

def total_after(num=0):
    """ return the total grains in the square """
    return long(sum([2 ** (a - 1) for a in range(1, num + 1)]))
