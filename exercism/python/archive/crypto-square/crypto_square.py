""" module to return the old-skool crypto square method for a given input """
# -*- coding:utf-8 -*-

from __future__ import unicode_literals

import math

def encode(src):
    """ Function to encode the given src string """
    if not src:
        return ''

    src = ''.join([x.lower() for x in src if x.isalpha() or x.isdigit()])

    rows = int(round(math.sqrt(len(src))))
    cols = int(math.ceil(len(src) / rows))

    if len(src) % rows != 0:
        cols += 1

    words = [src[i:i+cols] for i in range(0, len(src), cols)]

    return ' '.join([''.join(x[y] for x in words if y < len(x)) for y in range(cols)])
