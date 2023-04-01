"""Module to calculate the Collatz Conjecture of various numbers."""


def steps(number: int) -> int:
    """Calculate the number of steps in the Collatz Conjecture for the given number.

    :param number: int - input number to run through the Collatz Conjecture
    :return: int - number of steps through the Conjecture to reach 1.
    """

    if number <= 0:
        raise ValueError("Only positive integers are allowed")

    num_steps = 0
    while number != 1:
        num_steps += 1
        print(f"number: {number}, num_step: {num_steps}")

        if number % 2 == 0:
            number /= 2
        else:
            number = 3 * number + 1

    return num_steps
