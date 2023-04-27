"""Module to flatten the list of lists."""
# -*- coding: utf-8 -*-

from __future__ import unicode_literals


def helper(lister):
    """Generate the individual elements of the list of lists."""
    if hasattr(lister, "__iter__"):
        for each in lister:
            if hasattr(each, "__iter__"):
                for item in helper(each):
                    yield item
            else:
                yield each
    else:
        yield lister


def flatten(lister):
    """Return the flatter list of lists."""
    output = (item for each in lister for item in helper(each))

    return [x for x in output if x is not None]
