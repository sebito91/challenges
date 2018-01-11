""" Module to render perfect numbers """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

def is_perfect(num):
    """ function to check if perfect """
    if sum([1] + [each + (num // each) for each in xrange(2, int(num**0.5)+1) \
            if num % each == 0]) == num:
        return True

    return False
