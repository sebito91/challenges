""" Module to convert from legacy scoring system to new scoring system """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

def transform(old):
    """ Return the newly transformed scoring logic for Scrabble """
    return {each.lower(): key for key, val in old.iteritems() for each in val}

