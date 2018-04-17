""" Module to test the ACME Robot """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

NORTH, EAST, SOUTH, WEST = 0, 1, 2, 3

class Robot(object):
    """ Class Robot is our main instance of the robot """
    def __init__(self, bearing=NORTH, xPtr=0, yPtr=0):
        """ initialize our robot class """
        self.bearing = bearing
        self.xval, self.yval = xPtr, yPtr

    @property
    def coordinates(self):
        """ return the {x, y} coordinates of our robot """
        return (self.xval, self.yval)

    def turn_left(self):
        """ function to turn the robot left """
        self.bearing -= 1
        self.bearing %= 4

    def turn_right(self):
        """ function to turn the robot right """
        self.bearing += 1
        self.bearing %= 4

    def advance(self):
        """ function to move the robot forward """
        if self.bearing == NORTH:
            self.yval += 1

        if self.bearing == SOUTH:
            self.yval -= 1

        if self.bearing == EAST:
            self.xval += 1

        if self.bearing == WEST:
            self.xval -= 1

    def simulate(self, directions):
        """ function to move the robot a few times """
        if not directions:
            return

        for each in directions:
            if each == 'L':
                self.turn_left()
                continue

            if each == 'R':
                self.turn_right()
                continue

            if each == 'A':
                self.advance()
