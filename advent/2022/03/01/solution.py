#! /usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Module for 20221203, problem 01.

https://adventofcode.com/2022/day/3
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
    total = 0

    for line in data:
        splitsies = int(len(line) / 2)
        one, two = set(line[:splitsies]), set(line[splitsies:])
        overlap = f"{''.join(one & two)}"
        print(
            f"overlap: {overlap}, value: {ord(overlap) - 96}, len_one: {len(one)}, one: {one}, len_two: {len(two)}, two: {two}"
        )

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
