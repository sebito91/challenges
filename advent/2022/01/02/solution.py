#! /usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Module for 20221201, problem 01.

https://adventofcode.com/2022/day/1
"""

from collections import defaultdict
from pathlib import Path
from pprint import pprint


def new_elves_dict():
    """Return a new elves dict."""
    return {
        "total_cals": 0,
        "cals": [],
    }


def doit():
    """Run the game."""
    filename = Path("input")

    input_lines = []
    with filename.open(encoding="utf-8") as fh:
        input_lines = [x.strip() for x in fh.readlines()]

    print(f"num lines: {len(input_lines)}")
    print(f"some lines: {input_lines[:15]}")

    num_elves = 0
    elves_struct = defaultdict(new_elves_dict)
    for cals in input_lines:
        if not cals:
            num_elves += 1
            continue

        elves_struct[num_elves]["total_cals"] += int(cals)
        elves_struct[num_elves]["cals"].append(cals)

    elves_sorted = sorted(elves_struct.items(), key=lambda x: x[1]["total_cals"])
    print(f"elves: {sum([y['total_cals'] for x, y in elves_sorted[-3:]])}")


if __name__ == "__main__":
    doit()
