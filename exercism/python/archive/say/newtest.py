""" Module to convert given number to written number """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

def process_tens(num, ones, tens, teens, extra=False):
    """ helper function to process two-digit number """
    elems, add = str(num), ""

    if elems[0] == '0' and elems[1] == '0' or num == 0:
        return ""

    if extra:
        add = "and "

    if len(elems) == 1 or elems[0] == '0':
        return ''.join([add, ones[int(num)]])

    if elems[1] == '0':
        return ''.join([add, tens[int(elems[0])]])

    if elems[0] == '1' and elems[1] >= '4':
        return ''.join([add, ones[int(elems[1])], "teen"])

    if elems[0] == '1':
        return ''.join([add, teens[int(elems[1])]])

    return ''.join([add, tens[int(elems[0])], '-', ones[int(elems[1])]])

def process_hundreds(num, ones, tens, teens, extra):
    """ helper function to process hundreds """
    if not num:
        return ""

    elems = str(num)

    if elems[0] != '0' and len(elems) == 1:
        return ones[int(elems)]

    if elems[0] != '0' and len(elems) > 1:
        return ' '.join([ones[int(elems[0])], "hundred", process_tens(elems[1:], ones, tens, teens, extra)])

    return process_tens(elems[1:], ones, tens, teens, extra)

def process_thousands(num, ones, tens, teens, extra):
    """ helper function to process thousands """
    elems = str(num)

    lower, upper, ret = elems[len(elems)-3:], elems[:len(elems)-3], ""

    if int(lower) > 0:
        ret = process_hundreds(lower, ones, tens, teens, extra)

    if int(upper) == 0:
        return ret

    if ret:
        return ' '.join([process_hundreds(upper, ones, tens, teens, False), "thousand", ret])

    return ' '.join([process_hundreds(upper, ones, tens, teens, False), "thousand"])

def process_millions(num, ones, tens, teens, extra):
    """ helper function to process millions """
    elems = str(num)

    lower, upper, ret = elems[len(elems)-6:], elems[:len(elems)-6], ""
    print lower, upper, ret

    if len(upper) == 3:
        ret += process_hundreds(upper, ones, tens, teens, False) + " million "
    else:
        ret += process_tens(upper, ones, tens, teens, False) + " million "

    if int(lower):
        ret += process_thousands(lower, ones, tens, teens, extra)

    return ret

def process_billions(num, ones, tens, teens, extra):
    """ helper function to process billions """
    elems = str(num)

    lower, upper, ret = elems[len(elems)-9:], elems[:len(elems)-9], ""

    if len(upper) == 1:
        ret += ones[int(upper[0])] + " billion "
    else:
        ret += process_hundreds(upper, ones, tens, teens, False) + " billion "

    if int(lower):
        ret += process_millions(lower, ones, tens, teens, extra)

    return ret

def process_trillions(num, ones, tens, teens, extra):
    """ helper function to process trillions """
    elems = str(num)

    lower, upper, ret = elems[len(elems)-12:], elems[:len(elems)-12], ""

    if len(upper) == 1:
        ret += ones[int(upper[0])] + " trillion "
    else:
        ret += process_hundreds(upper, ones, tens, teens, False) + " trillion "

    if int(lower):
        ret += process_billions(lower, ones, tens, teens, extra)

    return ret

def say(num):
    """ function to convert num to written number """
    num = int(num)  # force 1e notation to int (eg. 1e9)

    if num < 0 or num >= 1000000000000000:
        raise AttributeError("malformed number passed: {}".format(num))

    thelen = len(str(num))

    ones_vals = [a for a in range(10)]

    ones_names = ["zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"]
    tens_names = ["", "ten", "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"]
    teens_names = ["ten", "eleven", "twelve", "thirteen"]

    ones = dict(zip(ones_vals, ones_names))
    tens = dict(zip(ones_vals, tens_names))
    teens = dict(zip(ones_vals, teens_names))

    if thelen == 1:
        return ones[num]

    if thelen < 3:
        return process_tens(num, ones, tens, teens, False)

    if thelen == 3:
        return process_hundreds(num, ones, tens, teens, True)

    if thelen > 3 and thelen < 7:
        return process_thousands(num, ones, tens, teens, True)

    if thelen > 6 and thelen < 10:
        return process_millions(num, ones, tens, teens, True)

    if thelen > 9 and thelen < 13:
        return process_billions(num, ones, tens, teens, True)

    if thelen > 12:
        return process_trillions(num, ones, tens, teens, True)

    raise AttributeError("Could not process number: {}".format(num))

def runit():
    """ run all the things """
    print "should be zero: {}".format(say(0))
    print "should be fourteen: {}".format(say(14))
    print "should be fifty: {}".format(say(50))

    print "should be one hundred: {}".format(say(100))
    print "should be ninety-eight: {}".format(say(98))
    #print "should be AttributeError: {}".format(say(-1))

    print "should be one thousand: {}".format(say(1000))
    print "should be one thousand and two: {}".format(say(1002))
    print "should be one thousand two hundred and thirty-four: {}".format(say(1234))
    print "should be one hundred twenty-three thousand four hundred and fifty-six: {}".format(say(123456))
    print "should be splits: {}".format(say(100000003))
    print "should be splits: {}".format(say(200234567890))
    print "should be splits: {}".format(say(111200234567890))
    print "should be one billion: {}".format(say(1e9))

if __name__ == "__main__":
    runit()
