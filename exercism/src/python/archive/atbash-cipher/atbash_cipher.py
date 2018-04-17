""" Module that implements the rudimentary atbash-cipher """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

def scramble(src):
    """ helper function to handle the scrambling """
    output = ""

    for each in src.lower():
        diff = ord(each) - ord('a')

        if diff >= 0 and diff < 26:
            output += chr(ord('a') + (25 - (ord(each) - ord('a'))))
        elif each >= '0' and each <= '9':
            output += each

    return output

def encode(src):
    """ Function to encode a src string into atbash-cipher """
    if not src:
        return None

    output = scramble(src)

    return ' '.join(output[i:i+5] for i in xrange(0, len(output), 5))

def decode(src):
    """ Function to decode a src string from atbash-cipher """
    if not src:
        return None

    return scramble(src)
