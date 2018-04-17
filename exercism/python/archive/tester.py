""" Module that counts the number of words in a phrase """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

def word_count_encode(src=''):
    """ Count the number of words in src, return dict with counts """
    src = src.strip().decode('utf-8')

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

def word_count_decode(src=''):
    """ Decode the string """
    src = src.strip().decode('utf-8')

    count = "0"
#    start = 0
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

        print "current: {}".format(output)
        index += 1

    return output

def runit():
    """ runit """
    print word_count_encode('yoyo')
    print word_count_encode('ahoy!')
    print word_count_encode('another.')
    print word_count_encode('last one?')
    print word_count_encode('')
    print word_count_encode('WWWBWWWBBBWWWB')
    print word_count_decode('3WB3W3B3WB')
    print word_count_decode(word_count_encode('WWWBWWWBBBWWWB'))

    print word_count_encode('AABBBCCCC')

if __name__ == "__main__":
    runit()
