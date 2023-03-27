"""Module to calculate the darts score."""


OUTER_RADIUS = 10
MIDDLE_RADIUS = 5
INNER_RADIUS = 1


def score(x: int, y: int) -> int:
    """Calculate the score given x- and y-axis values of the dart's position.

    :param x: int - the x-axis position of the dart.
    :param y: int - the y-axis position of the dart.
    :return: int - return the score given x- and y-axis position.

    Function to calculate the score for a game of darts given `x` and `y` score
    values. The concentric circles start at (0,0) and given their varying radii the score
    will adjust according to the following:

    1.  0 points - outside the outer circle
    2.  1 point  - inside the outer circle
    3.  5 points - inside the middle circle
    4. 10 points - inside the inner circle
    """

    # Using cartesian plane coordinates with a circle centered at (0, 0), our points
    # should be found using x^2 + y^2 = r^2.
    coordinates = pow(x, 2) + pow(y, 2)

    # outside the outer circle
    if coordinates > pow(OUTER_RADIUS, 2):
        return 0

    # inside the inner circle
    if coordinates <= pow(INNER_RADIUS, 2):
        return 10

    # inside the middle circle
    if coordinates <= pow(MIDDLE_RADIUS, 2):
        return 5

    # in the outer circle
    return 1
