""" Module to help convert long names to acronyms """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

import re

def abbreviate(src):
    """ function that receives a string into an acronym """
    if not src:
        return None

    words = (re.split(r'(?u)\W+', src.split(":")[0], flags=re.IGNORECASE))

    return ''.join([each[0].upper() for word in words for each in re.findall(r'(?u)[A-Za-z][^A-Z]*', word)])

def runit():
    """ run all the things """
    print "should be PNG: {}".format(abbreviate('Portable Network Graphics'))

    print "should be ROR: {}".format(abbreviate('Ruby on Rails'))

    print "should be HTML: {}".format(abbreviate('HyperText Markup Language'))

    print "should be FIFO: {}".format(abbreviate('First In, First Out'))

    print "should be PHP: {}".format(abbreviate('PHP: Hypertext Preprocessor'))

    print "should be CMOS: {}".format(abbreviate('Complementary metal-oxide semiconductor'))

if __name__ == "__main__":
    runit()
