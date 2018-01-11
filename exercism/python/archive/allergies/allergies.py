""" Module to provide the list of alleriges to which the user is affected """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

class Allergies(object):
    """ base class for our allergy set """
    def __init__(self, score):
        """ init function whereby we assign the user's allegies """
        self.bitwise_score = score % 256
        self.conditions = {
            0: "eggs",
            1: "peanuts",
            2: "shellfish",
            3: "strawberries",
            4: "tomatoes",
            5: "chocolate",
            6: "pollen",
            7: "cats",
        }

        self.ailments = []

        for each in range(8):
            if self.bitwise_score & (1 << each):
                self.ailments.append(self.conditions[each])

    def is_allergic_to(self, src=""):
        """ check if we're allergic to the given cause """
        if not src:
            return False

        for each in range(8):
            if self.conditions[each] != src:
                continue

            if self.bitwise_score & (1 << each):
                return True

        return False

    @property
    def lst(self):
        """ based on the score provided, return the list of ailments """
        return self.ailments
