""" Module to handle the pythangs """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals
import math
import fractions

def factors(n):
    """ function to generate factors """
    return [(i, n // i) for i in range(1, int(n**0.5) + 1) if n % i == 0]

def primitive_triplets(b):
    """ return the prim_trips """
    if b % 4 != 0:
        raise ValueError

    triples = set()
    for (n, m) in factors(b // 2):
        if (m - n) % 2 == 1 and fractions.gcd(m, n) == 1:
            triples.add(tuple(sorted((b, m**2 - n**2, m**2 + n**2))))
    return triples

def triplets_in_range(a, b):
    """ return the range of prim_trips """
    # Dickson's method
    triples = set()

    for s in range(1, b - 1):
        for t in range(s, b):
            r = int(math.sqrt(2 * s * t))
            if 2 * s * t == r**2:
                triple = (r + s, r + t, r + s + t)
                if all(a <= x <= b for x in triple):
                    triples.add(triple)
    return triples

def is_triplet(triplet):
    """ return whether it's prim_trip """
    a, b, c = sorted(triplet)
    return a**2 + b**2 == c**2
