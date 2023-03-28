"""Functions for tracking poker hands and assorted card tasks.

Python list documentation: https://docs.python.org/3/tutorial/datastructures.html
"""


def get_rounds(number: int) -> list:
    """Create a list containing the current and next two round numbers.

    :param number: int - current round number.
    :return: list - current round and the two that follow.
    """

    return [number + incr for incr in range(3)]


def concatenate_rounds(rounds_1: list, rounds_2: list) -> list:
    """Concatenate two lists of round numbers.

    :param rounds_1: list - first rounds played.
    :param rounds_2: list - second set of rounds played.
    :return: list - all rounds played.
    """

    return rounds_1 + rounds_2


def list_contains_round(rounds: list, number: int) -> bool:
    """Check if the list of rounds contains the specified number.

    :param rounds: list - rounds played.
    :param number: int - round number.
    :return: bool - was the round played?
    """

    return number in rounds


def card_average(hand: list) -> float:
    """Calculate and returns the average card value from the list.

    :param hand: list - cards in hand.
    :return: float - average value of the cards in the hand.
    """

    return sum(hand) / len(hand)


def approx_average_is_average(hand: list) -> bool:
    """Return if an average is using (first + last index values ) OR ('middle' card) == calculated average.

    :param hand: list - cards in hand.
    :return: bool - does one of the approximate averages equal the `true average`?
    """

    card_avg = card_average(hand)
    return card_avg == card_average([hand[0], hand[-1]]) or card_avg == hand[len(hand) // 2]


def average_even_is_average_odd(hand: list) -> bool:
    """Return if the (average of even indexed card values) == (average of odd indexed card values).

    :param hand: list - cards in hand.
    :return: bool - are even and odd averages equal?
    """

    return card_average(hand[0::2]) == card_average(hand[1::2])


def maybe_double_last(hand):
    """Multiply a Jack card value in the last index position by 2.

    :param hand: list - cards in hand.
    :return: list - hand with Jacks (if present) value doubled.
    """

    return hand[:-1] + [22 if hand[-1] == 11 else hand[-1]]
