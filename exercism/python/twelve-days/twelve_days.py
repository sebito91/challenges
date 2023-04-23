"""Module to recite the Twelve Days of Christmas song."""


DAYS = {
    1: "first",
    2: "second",
    3: "third",
    4: "fourth",
    5: "fifth",
    6: "sixth",
    7: "seventh",
    8: "eighth",
    9: "ninth",
    10: "tenth",
    11: "eleventh",
    12: "twelfth",
}


LYRICS = [
    "a Partridge in a Pear Tree",
    "two Turtle Doves",
    "three French Hens",
    "four Calling Birds",
    "five Gold Rings",
    "six Geese-a-Laying",
    "seven Swans-a-Swimming",
    "eight Maids-a-Milking",
    "nine Ladies Dancing",
    "ten Lords-a-Leaping",
    "eleven Pipers Piping",
    "twelve Drummers Drumming",
]


def recite(start_verse: int, end_verse: int) -> list[str]:
    """Generate the verses for the Twelve Days of Christmas song.

    :param start_verse: int - start verse number
    :param end_verse: int - end verse number
    :return: list[str] - series of verses for the song between `start_verse` and `end_verse`.
    """
    start_line = f"On the {DAYS[start_verse]} day of Christmas my true love gave to me: "

    for verse in range(start_verse - 1, 0, -1):
        start_line += f"{LYRICS[verse]}, "

    if start_verse != 1:
        start_line += "and "

    start_line += f"{LYRICS[0]}."
    if start_verse == end_verse:
        return [start_line]

    return [start_line] + recite(start_verse + 1, end_verse)
