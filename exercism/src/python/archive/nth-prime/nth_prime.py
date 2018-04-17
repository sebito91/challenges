""" Module to return the nth-prime """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

def nth_prime(num):
    """ return the prime at the number given, using the sieve """
    if num <= 0:
        raise ValueError("no zero-th prime")

    cap = 1000000   # 1M for fun
    sieve = [True] * (cap / 2)

    for val in xrange(3, int(cap ** 0.5) + 1, 2):
        if sieve[val / 2]:
            sieve[val * val / 2::val] = [False] * ((cap - val * val - 1) / (2 * val) + 1)

    deets = [2] + [2 * val + 1 for val in xrange(1, cap / 2) if sieve[val]]

    return deets[num-1]
