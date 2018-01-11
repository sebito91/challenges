""" Module to reply to given statements """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

def hey(src=''):
    """ function to reply to a given statement with some sass """
    src = src.strip()

    if not src:
        return "Fine. Be that way!"

    if src.isupper() or src[-1] == '!':
        return "Whoa, chill out!"

    if src[-1] == '?':
        return "Sure."

    return "Whatever."
