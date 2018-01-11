""" Module that's our tester """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

import re
import datetime

class MeetupDayException(Exception):
    """ our class to return a fault when encountered """
    def __init__(self, value):
        self.value = value

    def __str__(self):
        return repr(self.value)

def process_teenth(year, month, day):
    """ process teenth """
    wkday = datetime.date(year, month, 13).weekday()
    goal = day - wkday

    if goal < 0:
        goal += 7

    try:
        return datetime.date(year, month, 13 + goal)
    except ValueError:
        raise MeetupDayException("malformed teenth date")

def process_multiple(year, month, day, extra):
    """ process ranges like 1st, 2nd, 3rd, etc """
    wkday = datetime.date(year, month, 1).weekday()
    goal = day - wkday

    wk_mult = re.split(r'[a-z]+', extra)

    if goal < 0:
        goal += 7

    goal += 7 * (int(wk_mult[0]) - 1)

    try:
        return datetime.date(year, month, 1 + goal)
    except ValueError:
        raise MeetupDayException("date out of range for: {}".format(extra))

def process_last(year, month, day):
    """ process the last one """
    try:
        return process_multiple(year, month, day, '5th')
    except MeetupDayException:
        return process_multiple(year, month, day, '4th')
    else:
        raise MeetupDayException("could not find the last day of the month")

def printit(year, month, day, extra=''):
    """ function that tests our things out """
    if not (year or month or day):
        raise MeetupDayException("no date provided")

    days = {'Monday': 0, 'Tuesday': 1, 'Wednesday': 2, 'Thursday': 3, 'Friday': 4, 'Saturday': 5, 'Sunday': 6}

    if extra == 'teenth':
        return process_teenth(year, month, days[day])

    if extra == 'last':
        return process_last(year, month, days[day])

    return process_multiple(year, month, days[day], extra)

if __name__ == "__main__":
    print printit(2013, 5, 'Monday', 'teenth')
    print printit(2013, 5, 'Saturday', 'teenth')
    print printit(2013, 5, 'Tuesday', '1st')
    print printit(2013, 5, 'Tuesday', '2nd')
    print printit(2013, 5, 'Wednesday', '3rd')
    print printit(2013, 5, 'Tuesday', '4th')
    print printit(2013, 5, 'Thursday', 'last')
    print printit(2013, 3, 'Sunday', '4th')
