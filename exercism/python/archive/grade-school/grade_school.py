""" Module to handle school entries """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

class School(object):
    """ base object for our School instance """

    def __init__(self, name=None):
        """ initialize a new school """
        self._name = name or "Basic School"
        self._grades = dict(zip(range(1, 13), [[] for each in range(1, 13)]))

    def add(self, student=None, grade=0):
        """ helper function to add a student to a grade """
        if not student or grade == 0:
            raise ValueError("No student name and/or grade provided")

        try:
            self._grades[grade].append(student)
        except KeyError:
            raise ValueError("Incorrect grade given: {}".format(grade))

    def grade(self, grade=0):
        """ helper funnction to return the students in a grade """
        if grade == 0:
            raise ValueError("No grade provided")

        try:
            return set(sorted(self._grades[grade]))
        except KeyError:
            raise ValueError("Incorrect grade given: {}".format(grade))

    def sort(self):
        """ helper function to sort the dict """
        return [(key, tuple(sorted(self._grades[key]))) for key in sorted(self._grades.keys()) if self._grades[key]]
