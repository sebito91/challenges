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

TOLERANCE_CODE = {
    "grey": 0.05,
    "violet": 0.1,
    "blue": 0.25,
    "green": 0.5,
    "brown": 1,
    "red": 2,
    "gold": 5,
    "silver": 10
}


def resistor_label(colors: list[str]) -> str:
    """Generate the ohms from the given list of resistor colors.

    :param colors: list[str] - series of resistors to convert to ohms
    :return: str - number of ohms + tolerance for resistors
    """
    digits = [COLOR_CODE.get(color) for color in colors[:-1]]
    tolerance = f"±{TOLERANCE_CODE.get(colors[-1])}%"
    if len(digits) < 3:
        return "0 ohms"

    zero = "0"
    try:
        if len(digits) == 4:
            number = int(f"{digits[0]}{digits[1]}{digits[2]}{zero * (int(digits[3]))}")
        else:
            number = int(f"{digits[0]}{digits[1]}{zero * (int(digits[2]))}")
    except ValueError:
        number = 0

    if number < 1000:
        ohms = "ohms"
    elif number // 1_000 < 1_000:
        ohms = "kiloohms"
        number /= 1_000
    elif number // 1_000_000 < 1_000:
        ohms = "megaohms"
        number /= 1_000_000
    else:
        ohms = "gigaohms"
        number /= 1_000_000_000

    if f"{number}".split(".")[-1] == zero:
        number = int(number)

    return f"{number} {ohms} {tolerance}"
