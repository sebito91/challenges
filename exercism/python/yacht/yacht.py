"""Module to calculate the Yacht game score."""

from collections import defaultdict


# Score categories.
# Change the values as you see fit.

# multiply number of die value by that value (e.g. [3, 3, 3, 1, 2] = 3 * 3 * 3 * 3 = 27)
ONES = 1
TWOS = 2
THREES = 3
FOURS = 4
FIVES = 5
SIXES = 6

# total of all the dice if triple + pair
FULL_HOUSE = 7

# total of four matching numbers
FOUR_OF_A_KIND = 8

# dice in [1, 2, 3, 4, 5] score 30
LITTLE_STRAIGHT = 9

# dice in [2, 3, 4, 5, 6] score 30
BIG_STRAIGHT = 10

# sum of dice
CHOICE = 11

# five-of-a-kind scores 50
YACHT = 12


def score(dice: list[int], category) -> int:
    """Calculate the score for the given series of dice.

    :param dice: list[int] - series of dice to calculate the score
    :param category: X - specific score category
    :return: int - score of dice based on category
    """

    # prep a dict of the dice elements
    dice_dict = defaultdict(int)
    for die in dice:
        dice_dict[die] += 1

    keys = len(dice_dict.keys())

    if category in [ONES, TWOS, THREES, FOURS, FIVES, SIXES]:
        return category * dice_dict.get(category, 0)

    if category == FULL_HOUSE:
        if keys == 2 and sorted(dice_dict.values()) == [2, 3]:
            return sum(dice)

    if category == FOUR_OF_A_KIND:
        for key, count in dice_dict.items():
            if count >= 4:
                return 4 * key

    if category == LITTLE_STRAIGHT:
        if sorted(dice) == [1, 2, 3, 4, 5]:
            return 30

    if category == BIG_STRAIGHT:
        if sorted(dice) == [2, 3, 4, 5, 6]:
            return 30

    if category == CHOICE:
        return sum(dice)

    if category == YACHT:
        if keys == 1 and sorted(dice_dict.values()) == [5]:
            return 50

    return 0
