""" Module to handle clock input and changes """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

class Clock(object):
    """ The overarching class that will handle our clock functions """
    def __init__(self, hour, minute):
        """ initialize the clock """
        self._hour = 0
        self._minute = 0

        minutes = (hour * 60) + minute
        self.add(minutes)

    def add(self, minute):
        """ Add the requisite set of minutes to our time """
        one_day = 60 * 24
        total_mins = (self._minute + minute + (self._hour * 60)) % one_day

        self._hour = (total_mins / 60) % 24
        self._minute = total_mins % 60

        if self._hour < 0:
            self._hour += 24

        if self._minute < 0:
            self._minute += 60

            if self._hour == 0:
                self._hour = 23
            else:
                self._hour -= 1

        return "{:02}:{:02}".format(self._hour, self._minute)

    def __str__(self):
        """ How we return data for ourselves """
        return "{:02}:{:02}".format(self._hour, self._minute)

    def __eq__(self, other):
        return self.__dict__ == other.__dict__
