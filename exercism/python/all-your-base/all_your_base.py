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

    if not digits or all(digit == 0 for digit in digits):
        return [0]

    if not all((int(digit) < input_base and int(digit) >= 0) for digit in digits):
        raise ValueError("all digits must satisfy 0 <= d < input base")

    # convert to decimal
    input_sum = sum(digit[0] * pow(input_base, digit[1]) for digit in zip(digits[-1::-1], list(range(len(digits)))))

    if output_base < 2:
        raise ValueError("output base must be >= 2")

    # convert from decimal
    quotient, remainder = input_sum, 0
    output_digits = []

    while quotient != 0:
        quotient, remainder = quotient // output_base, quotient % output_base
        output_digits.append(remainder)

    return output_digits[-1::-1]
