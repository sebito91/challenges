"""Module to recite the Twelve Days of Christmas song."""


LYRICS = {
    1: ("first", "a Partridge in a Pear Tree"),
    2: ("second", "two Turtle Doves"),
    3: ("third", "three French Hens"),
    4: ("fourth", "four Calling Birds"),
    5: ("fifth", "five Gold Rings"),
    6: ("sixth", "six Geese-a-Laying"),
    7: ("seventh", "seven Swans-a-Swimming"),
    8: ("eighth", "eight Maids-a-Milking"),
    9: ("ninth", "nine Ladies Dancing"),
    10: ("tenth", "ten Lords-a-Leaping"),
    11: ("eleventh", "eleven Pipers Piping"),
    12: ("twelfth", "twelve Drummers Drumming"),
}


def recite(start_verse: int, end_verse: int) -> list[str]:
    """Generate the verses for the Twelve Days of Christmas song.

    :param start_verse: int - start verse number
    :param end_verse: int - end verse number
    :return: list[str] - series of verses for the song between `start_verse` and `end_verse`.
    """
    start_line = f"On the {LYRICS[start_verse][0]} day of Christmas my true love gave to me: "

    for verse in range(start_verse, 1, -1):
        start_line += f"{LYRICS[verse][1]}, "

    if start_verse != 1:
        start_line += "and "

    start_line += f"{LYRICS[1][1]}."
    if start_verse == end_verse:
        return [start_line]

    return [start_line] + recite(start_verse + 1, end_verse)
