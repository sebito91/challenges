""" Module to present the plants grown in the kindergarten_garden """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

class Garden(object):
    """ class to handle the plants in our kindergarten garden """
    def __init__(self, garden, students=None):
        """ our initialization function """
        if students:
            self._students = dict(zip(sorted(students), range(len(students))))
        else:
            self._students = dict(zip([
                "Alice", "Bob", "Charlie", "David",
                "Eve", "Fred", "Ginny", "Harriet",
                "Ileana", "Joseph", "Kincaid", "Larry"
            ], range(12)))

        self._veggies = {'R': 'Radishes', 'C': 'Clover', 'G': 'Grass', 'V': 'Violets'}

        self._raw_garden = garden
        self._row_one, self._row_two = self.generate()

    def generate(self):
        """ helper function to generate the garden """
        rows = self._raw_garden.split('\n')
        return [a for a in list(rows[0])], [b for b in list(rows[1])]

    def plants(self, student):
        """ property that returns the plants for a student """
        try:
            idx = self._students[student] * 2
        except KeyError:
            return []

        return [
            self._veggies[self._row_one[idx]],
            self._veggies[self._row_one[idx+1]],
            self._veggies[self._row_two[idx]],
            self._veggies[self._row_two[idx+1]]
        ]
