""" Module to handle the isogram function """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

def is_isogram(word):
    """ function to determine whether a word is an isogram (aka no repeats) """
    chars = set()

    for each in word.lower():
        if each in chars:
            return False

        if each.isalpha():
            chars.add(each)

    return True
