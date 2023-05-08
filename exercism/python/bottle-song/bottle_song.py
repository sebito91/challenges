"""Module to recite the Ten Green Bottles song."""

VALUES = {
    10: ("Ten", "bottles"),
    9: ("Nine", "bottles"),
    8: ("Eight", "bottles"),
    7: ("Seven", "bottles"),
    6: ("Six", "bottles"),
    5: ("Five", "bottles"),
    4: ("Four", "bottles"),
    3: ("Three", "bottles"),
    2: ("Two", "bottles"),
    1: ("One", "bottle"),
    0: ("No", "bottles"),
}


def recite(start: int, take: int = 1) -> list[str]:
    """Recite the Ten Green Bottles song using the provided `start` and `take` values.

    :param start: int - verse to start the song with
    :param take: int - number of verses to recite
    :return: list[str] - recital of verses taking into account the `start` and `take` values
    """
    if start > 10 or start < 1:
        raise ValueError(f"start value must be between 1 and 10, got: {start}")

    output = []
    for idx in range(take):
        verse_num = VALUES.get(start - idx)
        next_verse = VALUES.get(start - idx - 1)

        output.append(f"{verse_num[0]} green {verse_num[1]} hanging on the wall,")
        output.append(f"{verse_num[0]} green {verse_num[1]} hanging on the wall,")
        output.append("And if one green bottle should accidentally fall,")
        output.append(f"There'll be {next_verse[0].lower()} green {next_verse[1]} hanging on the wall.")

        if idx != (take - 1):
            output.append("")

    return output
