""" hello_world first function for exercism """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

def hello(name=''):
    """ function to return Hello and either name, or World """
    if not name:
        return "Hello, World!"
    else:
        return "Hello, {}!".format(name)
