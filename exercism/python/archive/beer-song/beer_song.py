""" Module to print 99 bottles song """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

def verse(num):
    """ print individual verse """
    verses = list()

    if num > 1:
        verses.append("{0} bottles of beer on the wall, {0} bottles of beer.".format(num))
    elif num == 1:
        verses.append("{0} bottle of beer on the wall, {0} bottle of beer.".format(num))
    else:
        verses.append("No more bottles of beer on the wall, no more bottles of beer.")

    if num > 2:
        verses.append("Take one down and pass it around, {} bottles of beer on the wall.\n".format(num - 1))
    elif num == 2:
        verses.append("Take one down and pass it around, {} bottle of beer on the wall.\n".format(num - 1))
    elif num == 1:
        verses.append("Take it down and pass it around, no more bottles of beer on the wall.\n")
    else:
        verses.append("Go to the store and buy some more, 99 bottles of beer on the wall.\n")

    return '\n'.join(verses)

def song(start, stop=0):
    """ print all the verses """
    return '\n'.join([verse(each) for each in xrange(start, stop-1, -1)] + [""])
