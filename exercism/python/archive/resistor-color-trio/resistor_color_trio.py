"""Module to translate resistor colors to ohms."""


COLOR_CODE = {
    "black": 0,
    "brown": 1,
    "red": 2,
    "orange": 3,
    "yellow": 4,
    "green": 5,
    "blue": 6,
    "violet": 7,
    "grey": 8,
    "white": 9
}


def label(colors: list[str]) -> str:
    """Generate the ohms from the given list of resistor colors.

    :param colors: list[str] - series of resistors to convert to ohms
    :return: str - number of ohms for resistors
    """
    digits = [COLOR_CODE.get(color) for color in colors]
    if len(digits) < 3:
        return ""

    zero = "0"
    try:
        number = int(f"{digits[0]}{digits[1]}{zero * (int(digits[2]))}")
    except ValueError:
        number = 0

    if number < 1000:
        ohms = "ohms"
    elif number // 1_000 < 1_000:
        ohms = "kiloohms"
        number //= 1_000
    elif number // 1_000_000 < 1_000:
        ohms = "megaohms"
        number //= 1_000_000
    else:
        ohms = "gigaohms"
        number //= 1_000_000_000

    return f"{number} {ohms}"
