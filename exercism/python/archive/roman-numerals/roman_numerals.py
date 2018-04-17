""" Module to convert arabic numbers into roman numerals """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

def numeral(num=0):
    """ function to convert arabic num into roman numeral """
    if num <= 0:
        raise ValueError("Invalid number provided")

    ones = dict(zip(range(11), ['', 'I', 'II', 'III', 'IV', 'V', 'VI', 'VII', 'VIII', 'IX', 'X']))
    tens = dict(zip(range(11), ['', 'X', 'XX', 'XXX', 'XL', 'L', 'LX', 'LXX', 'LXXX', 'XC', 'C']))
    hunds = dict(zip(range(11), ['', 'C', 'CC', 'CCC', 'CD', 'D', 'DC', 'DCC', 'DCCC', 'CM', 'M']))

    return ''.join(['M' * (num / 1000), hunds[(num / 100) % 10],
                    tens[(num % 100) / 10], ones[(num % 100) % 10]])
