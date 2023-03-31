"""Module to calculate the Raindrops module tests."""


def convert(number: int) -> str:
    """Calculate the Raindrops value for the provided number.

    :param number: int - number to convert from value to string, if divisible by any of 3, 5, or 7
    :return: str - return a string representing Raindrops if divisible by 3, 5, or 7; otherwise initial value
    """

    output = ""
    if number % 3 == 0:
        output += "Pling"

    if number % 5 == 0:
        output += "Plang"

    if number % 7 == 0:
        output += "Plong"

    if not output:
        return f"{number}"

    return output
