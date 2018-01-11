""" Module to add a 10^9 second to a given date """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals
import datetime

def add_gigasecond(src=''):
    """ function to add 10^9 seconds to a given datetime """
    return src + datetime.timedelta(seconds=1000000000)

