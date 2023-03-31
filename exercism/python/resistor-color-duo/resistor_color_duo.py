"""Module to return the color values for resistor bands."""


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


def value(colors: str) -> int:
    """Return the value of the provided resistor band colors.

    :param colors: str - dash-delimited list of colors for resistor bands.
    :return: int - value of color bands provided in `colors`
    """
    if len(colors) < 2:
        return COLOR_CODE.get(colors[0], -1)

    return int(f"{COLOR_CODE.get(colors[0], -1)}{COLOR_CODE.get(colors[1], -1)}")
