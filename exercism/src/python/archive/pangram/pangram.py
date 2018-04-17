""" Module to handle the pangram challenge """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

def is_pangram(string=''):
    """ Function to return whether we're a pangram, using every alphabet letter at least once """
    pangram = [1] * 26

    for each in string:
        char = ord(each.lower()) - ord('a')
        if char >= 0 and char < 26:
            pangram[char] = 0

    if 1 in pangram:
        return False

    return True
