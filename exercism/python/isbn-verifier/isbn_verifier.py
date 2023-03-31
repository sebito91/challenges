"""Module to verify a given ISBN-10 number."""


def is_valid(isbn: str) -> bool:
    """Check whether the provided ISBN-10 number is valid.

    :param isbn: str - ISBN-10 number to check for validity
    :return: bool - boolean value on whether provides ISBN-10 number is valid
    """

    isbn_sum = 0
    vals = list(zip(isbn.replace("-", ""), list(range(10, len(isbn) * -1, -1))))

    if len(vals) != 10:
        return False

    for number in vals:
        val, multiple = number[0], number[1]

        if val == "X" and multiple != 1:
            return False

        if val == "X":
            val = 10

        try:
            isbn_sum += int(val) * multiple
        except ValueError:
            return False

    return isbn_sum > 0 and isbn_sum % 11 == 0
