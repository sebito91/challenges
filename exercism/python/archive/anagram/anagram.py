""" Module to detect whether an anagram of a given string exists in a list
of possible outcomes """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

def check_items(items):
    """ helper function to check the list of pairs """
    for item in items:
        if item[0] != item[1]:
            return False

    return True

def detect_anagrams(src, outcomes=None):
    """ Given a src string, detect if there are anagrams presents in any
    of the list of outcomes provided """
    if not (src or outcomes):
        return []

    matches = []

    for each in outcomes:
        if len(each) != len(src):
            continue

        if each.lower() == src.lower():
            continue

        items = zip(sorted(src.lower()), sorted(each.lower()))
        if check_items(items):
            matches.append(each)

    return matches
