"""Module to convert a number, represented as a sequence of digits in one base, to any other base."""


def rebase(input_base: int, digits: list, output_base: int) -> list:
    """Rebase the given input value to a different base.

    :param input_base: int - initial base of the series of digits.
    :param digits: list - series of digits to convert from `input_base` to `output_base`.
    :param output_base: int - output base of the series of digits.
    :return: list - series of digits in the `output_base`

    Function to convert the given series of `digits` from `input_base` to `output_base`.
    """

    if input_base < 2:
        raise ValueError("input base must be >= 2")

    if output_base < 2:
        raise ValueError("output base must be >= 2")

    if sum(digits) == 0:
        return [0]

    if not all(0 <= digit < input_base for digit in digits):
        raise ValueError("all digits must satisfy 0 <= d < input base")

    # convert to decimal
    input_sum = sum(digit * pow(input_base, idx) for idx, digit in enumerate(digits[::-1]))

    # convert from decimal
    quotient, remainder = input_sum, 0
    output_digits = []

    while quotient != 0:
        quotient, remainder = divmod(quotient, output_base)
        output_digits.append(remainder)

    return output_digits[::-1]
