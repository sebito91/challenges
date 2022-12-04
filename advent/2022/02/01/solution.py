#! /usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Module for 20221202, problem 01.

https://adventofcode.com/2022/day/2
"""

from pathlib import Path


SCORES = {
    "A": 1,
    "B": 2,
    "C": 3,
    "X": 1,
    "Y": 2,
    "Z": 3,
}

POINTS = {
    "win": 6,
    "draw": 3,
    "lose": 0,
}


def _read_input(filename="input"):
    """Input filename."""
    filename = Path("input")
    with filename.open(encoding="utf-8") as fh:
        data = fh.readlines()

    return [line.strip() for line in data]


def calc_score(data=None):
    """Calculate the scores."""
    # A beats Z
    # B beats X
    # C beats Y
    total_score = 0
    for each in data:
        p1, p2 = each.split()
        total_score += SCORES[p2]
        print(f"     p1: {p1}, p2: {p2}, p2_score: {SCORES[p2]}, total_score: {total_score}")

        # draw case
        if SCORES[p2] == SCORES[p1]:
            total_score += POINTS["draw"]
            print(f"DRAW p1: {p1}, p2: {p2}, points: {POINTS['draw']}, total_score: {total_score}")
        elif p2 == "X" and p1 == "C":
            total_score += POINTS["win"]
            print(f"WIN  p1: {p1}, p2: {p2}, points: {POINTS['win']}, total_score: {total_score}")
        elif p2 == "Y" and p1 == "A":
            total_score += POINTS["win"]
            print(f"WIN  p1: {p1}, p2: {p2}, points: {POINTS['win']}, total_score: {total_score}")
        elif p2 == "Z" and p1 == "B":
            total_score += POINTS["win"]
            print(f"WIN  p1: {p1}, p2: {p2}, points: {POINTS['win']}, total_score: {total_score}")
        else:
            total_score += POINTS["lose"]
            print(f"LOSE p1: {p1}, p2: {p2}, points: {POINTS['lose']}, total_score: {total_score}")

    return total_score


def doit():
    """Return the correct answer."""
    data = _read_input()
    print(f"start of data: {data[:5]}")

    print(f"top score: {calc_score(data)}")


if __name__ == "__main__":
    doit()
