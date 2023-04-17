"""Module to recite The House That Jack Built."""

VERSES = [
    "",
    "the house that Jack built",
    "the malt",
    "the rat",
    "the cat",
    "the dog",
    "the cow with the crumpled horn",
    "the maiden all forlorn",
    "the man all tattered and torn",
    "the priest all shaven and shorn",
    "the rooster that crowed in the morn",
    "the farmer sowing his corn",
    "the horse and the hound and the horn",
]

ACTIONS = [
    "",
    "lay in",
    "ate",
    "killed",
    "worried",
    "tossed",
    "milked",
    "kissed",
    "married",
    "woke",
    "kept",
    "belonged to",
]


def gen_verse(verse_number: int) -> str:
    """Generate the given verse.

    :param verse_number: int - the verse to generate.
    :return: str - the given verse based on `verse_number`
    """
    initial_verse = f"This is {VERSES[verse_number]}"

    if verse_number == 1:
        return f"{initial_verse}."

    output = [initial_verse]
    for item in range(verse_number, 1, -1):
        output.append(f"that {ACTIONS[item - 1]} {VERSES[item - 1]}")

    return " ".join(output) + "."


def recite(start_verse: int, end_verse: int) -> str:
    """Recite the phrases between the start and end values.

    :param start_verse: int - start verse at line X
    :param end_verse: int - end verse at line Y
    :return: str - return the verses between the `start_verse` and `end_verse` values.
    """
    output = []
    for start in range(start_verse, end_verse + 1):
        output.append(gen_verse(start))

    return output
