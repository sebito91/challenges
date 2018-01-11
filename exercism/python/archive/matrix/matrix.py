""" Module to enable a matrix given inputs """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

class Matrix(object):
    """ definition of the base Matrix class """
    def __init__(self, lister=None):
        self._rows = [[int(a) for a in each.split(' ') if a.isnumeric()] \
                for each in lister.split('\n')]
        self._col_len = len(self._rows[0])
        self._row_len = len(self._rows)
        self._cols = [[self._rows[row][col] for row in xrange(0, self._row_len)] \
                for col in xrange(0, self._col_len)]

    @property
    def rows(self):
        """ return the rows from the Matrix """
        return self._rows

    @property
    def columns(self):
        """ return the columns from the Matrix """
        return self._cols
