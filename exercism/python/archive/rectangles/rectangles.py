"""Module to count the ASCII rectangles in a given input."""


def rectangles(strings: list[str]) -> int:
    """Count the number of ASCII rectangles within the given string input.

    :param strings: list[str] - series of ASCII rectangles to count
    :return: int - number of ASCII rectangles found in a given set of `strings` input.
    """
    output = set()

    hash_marks = sorted({(line_num, idx) for line_num, line in enumerate(strings) for idx, char in enumerate(line) if char == "+"})
    for idx, hash_mark in enumerate(hash_marks):
        trs = sorted({coord for coord in hash_marks[idx + 1:] if coord[0] == hash_mark[0]})
        brs = sorted({coord for coord in hash_marks[idx + 1:] for tr in trs if coord[1] == tr[1] and coord != tr})
        bls = sorted({coord for coord in hash_marks[idx + 1:] for br in brs if coord[0] == br[0] and coord[1] == hash_mark[1]})

        for tr in trs:
            for br in brs:
                for bl in bls:
                    if hash_mark[1] == bl[1] and br[0] == bl[0] and tr[1] == br[1]:
                        output.add((hash_mark, tr, br, bl))  # top-left, top-right, bottom-right, bottom-left

    bad_options = set()
    for option in output:
        tl, tr, br, bl = option

        # check top side, left-to-right
        chars = set(strings[tl[0]][tl[1]:tr[1]])
        if chars != {"+", "-"} and chars != {"+"}:
            bad_options.add(option)
            continue

        # check bottom side, left-to-right
        chars = set(strings[bl[0]][bl[1]:br[1]])
        if chars != {"+", "-"} and chars != {"+"}:
            bad_options.add(option)
            continue

        # check left side, top-to-bottom
        chars = {string[tl[1]] for string in strings[tl[0]:bl[0]]}
        if chars != {"+", "|"} and chars != {"+"}:
            bad_options.add(option)
            continue

        # check right side, top-to-bottom
        chars = {string[tr[1]] for string in strings[tr[0]:br[0]]}
        if chars != {"+", "|"} and chars != {"+"}:
            bad_options.add(option)
            continue

    return len(output - bad_options)
