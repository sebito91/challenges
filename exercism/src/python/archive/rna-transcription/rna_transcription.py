""" Module to run rna-transcription """
# -*- coding: utf-8 -*-

from __future__ import unicode_literals

def to_rna(dna=''):
    """ function to convert DNA to RNA strand """
    themap = {'G':'C', 'C':'G', 'T': 'A', 'A': 'U'}
    return ''.join([themap[each] for each in dna])
