""" Module to calculate the Luhn validity of a number """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

class Luhn(object):
    """ Class to handle the luhn validation of a given number """

    def __init__(self, number=0):
        """ instantiate the class for luhn validation """
        self._luhn = number

    def addends(self):
        """ Return list of the luhn elements """
        return [int(a) if (idx + 1) % 2 != 0 else \
                int(a) * 2 if int(a) < 5 else int(a) * 2 - 9 \
                for idx, a in enumerate(str(self._luhn)[::-1])]

    def checksum(self):
        """ Return the checksum of a given Luhn number """
        return sum(self.addends())

    def is_valid(self):
        """ Return whether the provided number is valid """
        return not bool(self.checksum() % 10)

    @staticmethod
    def create(number):
        """ Return newly created luhn number """
        if Luhn(number*10).is_valid():
            return number * 10

        return int(''.join([str(number), str(10 - (Luhn(number*10).checksum() % 10))]))
