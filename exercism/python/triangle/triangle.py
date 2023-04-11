"""Module to check for various triangle types."""


def check_lengths(sides: list[int]) -> bool:
    """Check whether the given sides are valid triangle lengths.

    :param sides: list[int] - the integer list of side lengths.
    :return: bool - whether sides are valid triangle lengths.
    """
    if len(sides) > 3 or any(side <= 0 for side in sides):
        return False

    if sides[0] + sides[1] < sides[2]:
        return False

    if sides[1] + sides[2] < sides[0]:
        return False

    if sides[0] + sides[2] < sides[1]:
        return False

    return True


def equilateral(sides: list[int]) -> bool:
    """Check whether the given sides represent an equilateral triangle.

    :param sides: list[int] - the integer list of side lengths.
    :return: bool - whether triangle is an equilateral
    """
    if not check_lengths(sides):
        return False

    return len(set(sides)) == 1


def isosceles(sides: list[int]) -> bool:
    """Check whether the given sides represent an isosceles triangle.

    :param sides: list[int] - the integer list of side lengths.
    :return: bool - whether triangle is an isosceles
    """
    if not check_lengths(sides):
        return False

    return 1 <= len(set(sides)) <= 2


def scalene(sides: list[int]) -> bool:
    """Check whether the given sides represent an scalene triangle.

    :param sides: list[int] - the integer list of side lengths.
    :return: bool - whether triangle is an scalene
    """
    if not check_lengths(sides):
        return False

    return len(set(sides)) == 3
