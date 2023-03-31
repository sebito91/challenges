"""Module to store and retrieve resistor colors."""


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


def color_code(color: str) -> int:
    """Retrieve the resistor color ID based on the provided `color`.

    :param color: str - color for which to retrieve the numeric ID
    :return: int - ID value for the given `color`
    """
    return COLOR_CODE.get(color, -1)


def colors() -> list[str]:
    """Return the list of resistor colors.

    :return: list[str] - the list of resistor color names
    """
    return list(COLOR_CODE.keys())
