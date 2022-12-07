#! /usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Module for 20221206, problem 02.

https://adventofcode.com/2022/day/6#part2
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
    elems = f"{data[0]}"
    print(f"elems: {elems[:5]}, len: {len(elems)}")
    marker_len = 14
    for x in range(len(elems) - marker_len):
        if len(set(elems[x: x + marker_len])) == marker_len:
            print(f"found it at char {x}, {elems[x:x + marker_len]}")
            return x + marker_len

    return 0


def doit():
    """Return the correct answer."""
    data = _read_input()
    print(f"start of data: {data[:5]}")

    print(f"top score: {calc_score(data)}")


if __name__ == "__main__":
    doit()
