"""Module to check for paired brackets."""


def is_paired(input_string: str) -> bool:
    """Check if the given `input_string` has matching parentheses/brackets.

    :param input_string: str - input string to check for matching brackets.
    :return: bool - boolean value representing whether brackets are matching.
    """
    next_to_close = []

    for char in input_string:
        if char in ["}", "]", ")"] and (len(next_to_close) == 0 or char != next_to_close[-1]):
            return False

        if char in ["}", "]", ")"] and char == next_to_close[-1]:
            next_to_close.pop(-1)

        if char == "{":
            next_to_close.append("}")

        if char == "[":
            next_to_close.append("]")

        if char == "(":
            next_to_close.append(")")

    return len(next_to_close) == 0
