""" Module that implements the rudimentary atbash-cipher """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

def scramble(src):
    """ helper function to handle the scrambling """

    output = ""

    for each in src.lower():
        diff = ord(each) - ord('a')

        if diff >= 0 and diff < 26:
            output += chr(ord('a') + (25 - (ord(each) - ord('a'))))
        elif each >= '0' and each <= '9':
            output += each

    return output

def encode(src):
    """ Function to encode a src string into atbash-cipher """
    if not src:
        return None

    output = scramble(src)

    return ' '.join(output[i:i+5] for i in xrange(0, len(output), 5))

def decode(src):
    """ Function to decode a src string from atbash-cipher """
    if not src:
        return None

    return scramble(src)

def runit():
    """ function to run all the things """
    print encode("no")

    print encode("yes")

    print encode("OMG")

    print encode("O M G")

    print encode("mindblowingly")

    print encode("Testing, 1 2 3, testing.")

    print encode("Truth is fiction.")

    print encode("The quick brown fox jumps over the lazy dog.")

    print decode("vcvix rhn")

    print decode("zmlyh gzxov rhlug vmzhg vkkrm thglm v")

if __name__ == "__main__":
    runit()
