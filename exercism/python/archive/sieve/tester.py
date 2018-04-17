
""" Module to handle the Sieve of Eratosthenes, prime numbers """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

def sieve(num):
    """ Given the num, return the primes up to the num """
    if num <= 0:
        return []

    thelist = [True] * (num + 1)

    for each in range(2, num + 1):
        item = each * 2

        while item < (num + 1):
            thelist[item] = False
            item += each

    return [a for a in range(2, num+1) if thelist[a]]

def runit():
    """ run all the things """
    print sieve(7)
    print sieve(10)
    print sieve(20)

if __name__ == "__main__":
    runit()
