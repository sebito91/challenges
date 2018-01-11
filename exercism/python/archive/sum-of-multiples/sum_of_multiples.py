""" Module to return the sum of multiples for a given list of inputs, up to a max number """
# -*- coding: utf-8 -*-

def sum_of_multiples(maxval, inputs):
    """ Function to calculate/return the sum of multiples up to maxval based on inputs """
    total = 0
    skippers = []

    items = (each for each in inputs if each > 0)

    for each in items:
        count = 1

        while (each * count) < maxval:
            if each * count in skippers:
                count += 1
                continue

            total += each * count
            skippers.append(each * count)

            count += 1

    return total
