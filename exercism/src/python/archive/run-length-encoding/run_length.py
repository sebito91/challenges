""" Module to enable run-length encoding (RLE) and decoding (RLD) """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

def decode_src(src):
    """ helper to decode stuffs """
    try:
        return src.decode('utf-8')
    except UnicodeEncodeError:
        return src

def encode(src=''):
    """ function to handle encoding """
    #src = src.strip().decode('utf-8')
    src = decode_src(src)

    count = 0
    start = 0
    index = 0
    output = ""

    while index < len(src):
        if src[start] == src[index]:
            count += 1
            index += 1
        else:
            if count > 1:
                output += "{}{}".format(count, src[start])
            else:
                output += "{}".format(src[start])

            start = index
            count = 0

    if start > 0 and src[start] != src[start-1]:
        if count > 1:
            output += "{}{}".format(count, src[start])
        else:
            output += "{}".format(src[start])

    return output


def decode(src=''):
    """ function to handle decoding """
    #src = src.strip().decode('utf-8')
    src = decode_src(src)

    count = "0"
    index = 0
    output = ""

    while index < len(src):
        if src[index].isdigit():
            count += src[index]
            index += 1
            continue

        if int(count) > 1:
            output += "{}".format(src[index] * int(count))
            count = "0"
        else:
            output += src[index]

        index += 1

    return output

# REALLY SLICK ANSWER
#import itertools
#import re
#
#
#def encode(x):
#    counts = [(len(list(g)), c) for c, g in itertools.groupby(x)]
#    return ''.join(str(e) for cs in counts for e in cs if e != 1)
#
#
#def decode(x):
#    return re.sub(r'(\d+)(\D)', lambda m: m.group(2) * int(m.group(1)), x)
