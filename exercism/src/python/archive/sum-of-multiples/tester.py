""" Module to return the sum of multiples for a given list of inputs, up to a max number """
# -*- coding: utf-8 -*-

def sum_of_multiples(maxval, inputs):
    """ Function to calculate/return the sum of multiples up to maxval based on inputs """
    total = []
    count = 1
    skippers = []

    items = (each for each in inputs if each > 0)

    for each in items:
        count = 1

        while (each * count) < maxval:
            if each * count in skippers:
                count += 1
                continue

            total.append(each * count)
            skippers.append(each * count)

            count += 1

    return sum(total)

def runit():
    """ run all the things """
    print "should be 0: {}".format(sum_of_multiples(1, [3, 5]))

    print "should be 3: {}".format(sum_of_multiples(4, [3, 5]))

    print "should be 23: {}".format(sum_of_multiples(10, [3, 5]))

    print "should be 2318: {}".format(sum_of_multiples(100, [3, 5]))

    print "should be 233168: {}".format(sum_of_multiples(1000, [3, 5]))

    print "should be 51: {}".format(sum_of_multiples(20, [7, 13, 17]))

    print "should be 30: {}".format(sum_of_multiples(15, [4, 6]))

    print "should be 4419: {}".format(sum_of_multiples(150, [5, 6, 8]))

    print "should be 2203160: {}".format(sum_of_multiples(10000, [43, 47]))

    print "should be 0: {}".format(sum_of_multiples(10, [0]))

    print "should be 45: {}".format(sum_of_multiples(10, [0, 1]))


if __name__ == "__main__":
    runit()
