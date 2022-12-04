#! /usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Module for 20221203, problem 02.

https://adventofcode.com/2022/day/3#part2
"""

from collections import defaultdict
from functools import reduce
from pathlib import Path


def _read_input(filename="input"):
    """Input filename."""
    filename = Path("input")
    with filename.open(encoding="utf-8") as fh:
        data = fh.readlines()

    return [line.strip() for line in data]


def _dict_item():
    """Return the dict item."""
    return {
        "rucks": [],
        "badge": None,
    }


def calc_score(data=None):
    """Calculate the score of priority items."""
    total = 0

    groups = defaultdict(_dict_item)

    for idx, line in enumerate(data):
        groups[int(idx / 3)]["rucks"].append(set(line))

    for idx, group in groups.items():
        overlap = f"{''.join(reduce(lambda x, y: x & y, (x for x in group['rucks'])))}"

        total += ord(overlap.lower()) - 96
        if overlap.isupper():
            print(
                f"overlap is upper: {overlap}, lower: {ord(overlap.lower()) - 96}, upper: {ord(overlap.lower()) - 96 + 27}"
            )
            total += 26

    return total


def doit():
    """Return the correct answer."""
    data = _read_input()
    print(f"start of data: {data[:5]}")

    print(f"top score: {calc_score(data)}")


if __name__ == "__main__":
    doit()
