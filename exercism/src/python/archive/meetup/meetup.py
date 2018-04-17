""" Module to handle meetups """
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
    """ given a date, provide the official date in the teens for that month """
    wkday = datetime.date(year, month, 13).weekday()
    goal = day - wkday

    if goal < 0:
        goal += 7

    try:
        return datetime.date(year, month, 13 + goal)
    except ValueError:
        raise MeetupDayException("malformed teenth date")

def process_multiple(year, month, day, extra):
    """ given a date, process the 'multiple' of the given date for that month """
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
    """ given a date, tell us the last occurrence of that date in a the given month """
    try:
        return process_multiple(year, month, day, '5th')
    except MeetupDayException:
        return process_multiple(year, month, day, '4th')
    else:
        raise MeetupDayException("could not find the last day of the month")

def meetup_day(year, month, day, extra=''):
    """ meetup_day takes a given date and returns the official meetup timeframe """
    if not (year or month or day):
        raise MeetupDayException("no date provided")

    days = {'Monday': 0, 'Tuesday': 1, 'Wednesday': 2, 'Thursday': 3, 'Friday': 4, 'Saturday': 5, 'Sunday': 6}

    if extra == 'teenth':
        return process_teenth(year, month, days[day])

    if extra == 'last':
        return process_last(year, month, days[day])

    return process_multiple(year, month, days[day], extra)
