#! /usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Module for 20221204, problem 02.

https://adventofcode.com/2022/day/4#part2
"""

from pathlib import Path


def _read_input(filename="input"):
    """Input filename."""
    filename = Path("input")
    with filename.open(encoding="utf-8") as fh:
        data = fh.readlines()

    return [line.strip() for line in data]


def calc_score(data=None):
    """Calculate the score of priority items."""
    total_overlaps = 0

    for line in data:
        a_lower, mid, b_upper = line.split("-")
        a_upper, b_lower = mid.split(",")

        a_lower = int(a_lower)
        a_upper = int(a_upper)
        b_lower = int(b_lower)
        b_upper = int(b_upper)

        print(
            f"line: {line}, a_lower: {a_lower}, a_upper: {a_upper}, b_lower: {b_lower}, b_upper: {b_upper}"
        )

        if a_lower <= b_upper and b_lower <= a_upper:
            print(f"a in b, would add line: {line}")
            total_overlaps += 1
        elif b_lower <= a_upper and a_lower <= b_upper:
            print(f"b in a, would add line: {line}")
            total_overlaps += 1

    return total_overlaps


def doit():
    """Return the correct answer."""
    data = _read_input()
    print(f"start of data: {data[:5]}")

    print(f"top score: {calc_score(data)}")


if __name__ == "__main__":
    doit()
