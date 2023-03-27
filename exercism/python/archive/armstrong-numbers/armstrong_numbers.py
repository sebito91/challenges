"""Module to determine whehter a given number is an Armstrong number."""


def is_armstrong_number(number: int) -> bool:
    """Determine whether the provided `number` is an Armstrong number.

    :param number: int - the value to check whether if is an Armstrong number.
    :return: bool - boolean value on whether provided `number` is an Armstrong number.

    Function to determine whether the given `number` is an Armstrong number; the equivalent of
    the sum of each digit raised to the number of digits in the number.
    """

    str_number = f"{number}"
    return sum(pow(int(x), len(str_number)) for x in str_number) == number
