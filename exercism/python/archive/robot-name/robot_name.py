""" Module to determine a random name for the robot """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals
from datetime import datetime

import random
import string

class Robot(object):
    """ base class for our robot """

    def __init__(self):
        self.reset()

    @property
    def name(self):
        """ return the robot's name """
        return self._name

    @staticmethod
    def gen_name():
        """ generate a unique name """
        random.seed(datetime.now())

        letters = [random.choice(string.letters.upper()) for x in xrange(0, 2)]
        numbers = [str(random.randint(0, 9)) for x in xrange(0, 3)]

        return ''.join(letters + numbers)

    def reset(self):
        """ reset the name """
        self._name = self.gen_name()
