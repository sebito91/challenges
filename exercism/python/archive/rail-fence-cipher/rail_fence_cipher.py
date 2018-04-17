""" Module to implement the rail-fence-cipher """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals
from collections import defaultdict

def prep_rows(matrix, msg_len):
    """ helper to print the rows """
    num_pass = max([len(value) for value in matrix.values()])

    lister = [matrix[0].pop(0)]
    for idx in xrange(0, num_pass):
        if idx % 2 == 0:
            lister.append(''.join([matrix[k].pop(0) for k in sorted(matrix.keys())[1:] \
                    if len(matrix[k]) > 0]))
        else:
            lister.append(''.join([matrix[k].pop(0) for k in reversed(sorted(matrix.keys())[:-1]) \
                    if len(matrix[k]) > 0]))

    return ''.join(lister)

def encode(msg, rails):
    """ encodes the given message """
    matrix = defaultdict(list)

    matrix[0] = [msg[item] for item in xrange(0, len(msg), (2 * (rails - 1)))]
    matrix[rails - 1] = [msg[item] for item in xrange(rails - 1, len(msg), (2 * (rails - 1)))]
    count = elems = len(matrix[0])

    for each in xrange(1, rails - 1):
        for item in xrange(each, len(msg), (2 * (rails - 1))):
            matrix[each].append(msg[item])
            if (item + (2 * (rails - 1 - each))) < len(msg):
                matrix[each].append(msg[item + (2 * (rails - 1 - each))])

    return ''.join([''.join(matrix[k]) for k in sorted(matrix.keys())])

def decode(msg, rails):
    """ decodes the given message """
    matrix = defaultdict(list)

    matrix[0] = [msg[item] for item in xrange(0, len(msg)) if item * (2 * (rails - 1)) < len(msg)]
    count = elems = len(matrix[0])

    if ((len(msg) - 1) - ((elems - 1) * (2 * (rails - 1)))) > 2:
        offset = 0
    else:
        offset = 2 - ((len(msg) - 1) - ((elems - 1) * (2 * (rails - 1))))

    for each in xrange(1, rails - 1):
        for item in xrange(0, 2 * elems - offset):
            matrix[each].append(msg[count])
            count += 1

    matrix[rails-1] = list(msg[count:])
    return prep_rows(matrix, len(msg))
