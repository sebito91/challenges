""" Module to implement a rotational-cipher """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

def rotate(msg, offset):
    """ function to handle the rotational cipher """

    lister = []
    for each in msg:
        calc = ord(each) + (offset % 26)

        if not each.isalpha():
            lister.append(each)
        elif calc > ord('z'):
            lister.append(chr(ord('a') + (calc - ord('z') - 1)))
        elif calc > ord('Z') and ord(each) < ord('a'):
            lister.append(chr(ord('A') + (calc - ord('Z') - 1)))
        else:
            lister.append(chr(calc))

    return ''.join(lister)
