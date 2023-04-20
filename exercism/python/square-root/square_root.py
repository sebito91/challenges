"""Module to calculate the square root of a given number."""


def square_root(number: int) -> int:
    """Calculate the square root of the given number without using the stdlib.

    :param number: int - replicand for root
    :return: int - square root of the replicand
    """
    # using Heron's Method we need to chose an arbitrarily close number to start
    # https://en.wikipedia.org/wiki/Methods_of_computing_square_roots#Heron's_method
    initial_value = number * 0.1

    # from the `initial_value` we now work to converge to the square root. Once two
    # iterations of the result are the same, we have our answer
    root, last_value = initial_value, 0

    while root != last_value:
        last_value = root
        root = 0.5 * (root + (number / root))

    return root
