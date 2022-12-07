#! /usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Module for 20221205, problem 02.

https://adventofcode.com/2022/day/5#part2
"""

from pathlib import Path
from pprint import pprint


def _read_input(filename="input"):
    """Input filename."""
    filename = Path("input")
    with filename.open(encoding="utf-8") as fh:
        data = fh.readlines()

    return [line.strip() for line in data]


def calc_score(data=None):
    """Calculate the score of priority items."""
    # these are manually added from the input
    piles = {
        1: ["R", "S", "L", "F", "Q"],
        2: ["N", "Z", "Q", "G", "P", "T"],
        3: ["S", "M", "Q", "B"],
        4: ["T", "G", "Z", "J", "H", "C", "B", "Q"],
        5: ["P", "H", "M", "B", "N", "F", "S"],
        6: ["P", "C", "Q", "N", "S", "L", "V", "G"],
        7: ["W", "C", "F"],
        8: ["Q", "H", "G", "Z", "W", "V", "P", "M"],
        9: ["G", "Z", "D", "L", "C", "N", "R"],
    }

    for line in [line for line in data if line.startswith("move")]:
        elems = line.split()
        count, src, dst = int(elems[1]), int(elems[3]), int(elems[5])

        print(
            f"line: {line}, count: {count}, src: {src} {piles[src][-count:]}, dst: {dst}"
        )
        to_move = piles[src][-count:]
        #        to_move.reverse()

        piles[dst] += to_move
        del piles[src][-count:]
        print(
            f"   move: {piles[src][-count:]}, to_move: {to_move}, dst: {dst} {piles[dst]}"
        )

    return "".join(pile[-1] for x, pile in piles.items())


def doit():
    """Return the correct answer."""
    data = _read_input()
    print(f"start of data: {data[:5]}")

    print(f"top score: {calc_score(data)}")


if __name__ == "__main__":
    doit()
