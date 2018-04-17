""" module to flatten a given array into its sub-elements """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

def actually_flatten(array):
    """ helper function to flatten the given array """


def flatten(array):
    """ function flatten our array """
    return [x for x in actually_flatten(array) if x != [] and x is not None]

