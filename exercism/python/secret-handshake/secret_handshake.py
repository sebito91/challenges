# -*- coding: utf-8 -*-
"""Module to return the secret handshake."""

from __future__ import unicode_literals

_shakes = {
    1: "wink",
    10: "double blink",
    100: "close your eyes",
    1000: "jump",
    }


def handshake(num):
    """Given a decimal input, return the sequence of stuff."""
    print "DEBUG -- num is: {}".format(num)


def code(seq):
    """Given a sequence of stuff, return the code."""
    test = zip(_shakes.values(), _shakes.keys())

    print "DEBUG -- code is: {}, test: {}".format(seq, test)
