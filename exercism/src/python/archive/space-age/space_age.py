""" Module to return the 'earth' time on different planets """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

class SpaceAge(object):
    """ structure to handle our time conversion """

    def __init__(self, seconds=0):
        """ Instantiate our space_age calculator """
        self._seconds = float(seconds)
        self._earth = 31557600.0

    @property
    def seconds(self):
        """ return the seconds we've stored at instantiation """
        return self._seconds

    def on_earth(self):
        """ Return time on earth """
        return float("{:.2f}".format(self._seconds / self._earth))

    def on_mercury(self):
        """ Return time on mercury """
        return float("{:.2f}".format(self._seconds / (self._earth * 0.2408467)))

    def on_venus(self):
        """ Return time on venus """
        return float("{:.2f}".format(self._seconds / (self._earth * 0.61519726)))

    def on_mars(self):
        """ Return time on mars """
        return float("{:.2f}".format(self._seconds / (self._earth * 1.8808158)))

    def on_jupiter(self):
        """ Return time on jupiter """
        return float("{:.2f}".format(self._seconds / (self._earth * 11.862615)))

    def on_saturn(self):
        """ Return time on saturn """
        return float("{:.2f}".format(self._seconds / (self._earth * 29.447498)))

    def on_uranus(self):
        """ Return time on uranus """
        return float("{:.2f}".format(self._seconds / (self._earth * 84.016846)))

    def on_neptune(self):
        """ Return time on neptune """
        return float("{:.2f}".format(self._seconds / (self._earth * 164.79132)))
