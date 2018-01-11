""" Module to test out saddle points in a matrix """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

def saddle_points(mat):
    """ function to determine saddle points in a given matrix """
    if len(mat) <= 0:
        return set()

    if len(set([len(each) for each in mat])) > 1:
        raise ValueError("malformed matrix")

    inv = zip(*mat)
    return set([(a, b) for a in xrange(0, len(mat)) for b in xrange(0, len(mat[0])) \
            if mat[a][b] >= max(mat[a]) and mat[a][b] <= min(inv[b])])
