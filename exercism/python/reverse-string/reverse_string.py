"""Module to test out reversing a string."""


def reverse(text: str) -> str:
    """Revers the provided text.

    :param text: str - input value to reverse
    :return: str - reversed version of the input `text`
    """
    return f"{text[::-1]}"
