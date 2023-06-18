"""Module to select the best poker hand."""

from collections import defaultdict


# card hands in priority order - https://en.wikipedia.org/wiki/List_of_poker_hands
FIVE_OF_A_KIND = 10
STRAIGHT_FLUSH = 9
FOUR_OF_A_KIND = 8
FULL_HOUSE = 7
FLUSH = 6
STRAIGHT = 5
THREE_OF_A_KIND = 4
TWO_PAIR = 3
ONE_PAIR = 2
HIGH_CARD = 1

CARDS = {
    "2": 2,
    "3": 3,
    "4": 4,
    "5": 5,
    "6": 6,
    "7": 7,
    "8": 8,
    "9": 9,
    "10": 10,
    "J": 11,
    "Q": 12,
    "K": 13,
    "A": 14
}


def best_hands(hands: list[str]) -> list[str]:
    """Choose the best poker hand from the set provided.

    :param hands: list[str] - series of poker hands
    :return: list[str] - best poker hand from set of `hands`.
    """
    output = process_hands(hands)
    return [option["hand"] for option in output[max(output.keys())]]


def create_cards(hand: list[str], flip_ace: bool = False) -> list[dict]:
    """Create the set of cards to process.

    :param hand: list[str] - hand of cards to process
    :param flip_ace: bool - check for an ace being low-card
    :return: list[dict] - individual cards broken into "num" and "suit" components.
    """
    try:
        cards = sorted([{"num": CARDS.get(card[:-1]), "suit": card[-1]} for card in hand.split()], key=lambda x: x["num"])

        if flip_ace:
            cards = sorted([{"num": card["num"], "suit": card["suit"]} if card["num"] != CARDS.get("A") else {"num": 1, "suit": card["suit"]} for card in cards if card], key=lambda x: x["num"])
    except ValueError as exc:
        raise ValueError(f"need five cards per hand, got {len(hand.split())}") from exc

    return cards


def generate_num_dict(hand: list[str]) -> dict:
    """Genereate a dictionary of card numbers and suits.

    :param hand: list[str] - hand to parse into a dictionary
    :return: dict - dictionary of card numbers and their suit.
    """
    nums_dict = defaultdict(list[str])

    for card in hand.split():
        nums_dict[CARDS.get(card[:-1])].append(card[-1])

    return nums_dict


def generate_suit_dict(hand: list[str]) -> dict:
    """Genereate a dictionary of card suits and their numbers.

    :param hand: list[str] - hand to parse into a dictionary
    :return: dict - dictionary of card numbers and their suit.
    """
    suit_dict = defaultdict(list)

    for card in hand.split():
        suit_dict[card[-1]].append(CARDS.get(card[:-1]))

    return suit_dict


def process_hands(hands: list[str]) -> list[int]:
    """Process the given hand into one of the set of valid hands.

    :param hands: list[str] - set of hands of cards to process
    :return: list[int] - set of hands categorized by priority
    """
    output = defaultdict(list[int])
    flip_ace = True

    for hand in hands:
        nums_dict = defaultdict(list)
        suit_dict = defaultdict(list)
        card_type = None

        cards = create_cards(hand)
        check_for_ace = cards
        if any(card["num"] == CARDS.get("A") for card in cards):
            check_for_ace = create_cards(hand, flip_ace)

        for card in hand.split():
            nums_dict[CARDS.get(card[:-1])].append(card[-1])
            suit_dict[card[-1]].append(CARDS.get(card[:-1]))

        if all(card["num"] == cards[0]["num"] for card in cards):
            card_type = FIVE_OF_A_KIND
            output[card_type] += [{"cards": cards, "hand": hand}]

            if len(output[card_type]) > 1:
                max_hands = [max(card["num"] for card in create_cards(option["cards"])) for option in output[card_type]]
                max_of_max_hands = max(max_hands)
                output[card_type] = [output[card_type][idx] for idx, max_hand in enumerate(max_hands) if max_hand == max_of_max_hands]

            if card_type == FIVE_OF_A_KIND:
                continue

        if all(card["suit"] == cards[0]["suit"] for card in cards):
            if all(cards[0]["num"] + idx == card["num"] for idx, card in enumerate(cards)):
                card_type = STRAIGHT_FLUSH
                output[card_type] += [{"cards": cards, "hand": hand}]
            elif all(check_for_ace[0]["num"] + idx == card["num"] for idx, card in enumerate(check_for_ace)):
                card_type = STRAIGHT_FLUSH
                output[card_type] += [{"cards": check_for_ace, "hand": hand}]

            if card_type == STRAIGHT_FLUSH and len(output[card_type]) > 1:
                # check each hand
                max_hands = [max(card["num"] for card in option["cards"]) for option in output[card_type]]
                max_of_max_hands = max(max_hands)
                possible_hands = [output[card_type][idx] for idx, max_hand in enumerate(max_hands) if max_hand == max_of_max_hands]

                if len(possible_hands) > 1:
                    max_hands = [max(option["cards"][-1]["num"] for option in output[card_type])]
                    max_of_max_hands = max(max_hands)
                    output[card_type] = [possible_hands[idx] for idx, max_hand in enumerate(max_hands) if max_hand == max_of_max_hands]
                else:
                    output[card_type] = possible_hands

            if card_type == STRAIGHT_FLUSH:
                continue

        if any(len(suits) == 4 for num, suits in nums_dict.items()):
            card_type = FOUR_OF_A_KIND
            output[card_type] += [{"cards": cards, "hand": hand}]

            if len(output[card_type]) > 1:
                max_hands = [max(num for num, suits in generate_num_dict(option["hand"]).items() if len(suits) == 4) for option in output[card_type]]
                max_of_max_hands = max(max_hands)
                possible_hands = [output[card_type][idx] for idx, max_hand in enumerate(max_hands) if max_hand == max_of_max_hands]

                if len(possible_hands) > 1:
                    max_hands = [max(card["num"] for card in option["cards"]) for option in possible_hands]
                    max_of_max_hands = max(max_hands)
                    output[card_type] = [possible_hands[idx] for idx, max_hand in enumerate(max_hands) if max_hand == max_of_max_hands]
                else:
                    output[card_type] = possible_hands

            if card_type == FOUR_OF_A_KIND:
                continue

        if any(len(suits) == 3 for num, suits in nums_dict.items()) and any(len(suits) == 2 for num, suits in nums_dict.items()):
            card_type = FULL_HOUSE
            output[card_type] += [{"cards": cards, "hand": hand}]

            if len(output[card_type]) > 1:
                max_hands = [max(num for num, suits in generate_num_dict(option["hand"]).items() if len(suits) == 3) for option in output[card_type]]
                max_of_max_hands = max(max_hands)
                possible_hands = [output[card_type][idx] for idx, max_hand in enumerate(max_hands) if max_hand == max_of_max_hands]

                if len(possible_hands) > 1:
                    max_hands = [max(card["num"] for card in option["cards"]) for option in possible_hands]
                    max_hands = [max(num for num, suits in generate_num_dict(option["hand"]).items() if len(suits) == 2) for option in possible_hands]
                    max_of_max_hands = max(max_hands)
                    output[card_type] = [possible_hands[idx] for idx, max_hand in enumerate(max_hands) if max_hand == max_of_max_hands]
                else:
                    output[card_type] = possible_hands

            if card_type == FULL_HOUSE:
                continue

        if len(suit_dict.keys()) == 1:
            card_type = FLUSH
            output[card_type] += [{"cards": cards, "hand": hand}]

            if len(output[card_type]) > 1:
                cards_dict = defaultdict(list)
                temp_output = defaultdict(list[str])

                for option in output[card_type]:
                    for card in option["cards"]:
                        cards_dict[card["num"]].append(option)

                for card_num, options in cards_dict.items():
                    if len(options) > 1:
                        continue

                    temp_output[card_num] += options

                if temp_output:
                    output[card_type] = temp_output[max(temp_output.keys())]

            if card_type == FLUSH:
                continue

        if all(cards[0]["num"] + idx == card["num"] for idx, card in enumerate(cards)) or all(check_for_ace[0]["num"] + idx == card["num"] for idx, card in enumerate(check_for_ace)):
            if all(cards[0]["num"] + idx == card["num"] for idx, card in enumerate(cards)):
                card_type = STRAIGHT
                output[card_type] += [{"cards": cards, "hand": hand}]
            elif all(check_for_ace[0]["num"] + idx == card["num"] for idx, card in enumerate(check_for_ace)):
                card_type = STRAIGHT
                output[card_type] += [{"cards": check_for_ace, "hand": hand}]

            if card_type == STRAIGHT and len(output[card_type]) > 1:
                max_hands = [max(card["num"] for card in option["cards"]) for option in output[card_type]]
                max_of_max_hands = max(max_hands)
                possible_hands = [output[card_type][idx] for idx, max_hand in enumerate(max_hands) if max_hand == max_of_max_hands]

                if len(possible_hands) > 1:
                    max_hands = [max(option["cards"][-1]["num"] for option in output[card_type])]
                    max_of_max_hands = max(max_hands)
                    output[card_type] = [possible_hands[idx] for idx, max_hand in enumerate(max_hands) if max_hand == max_of_max_hands]
                else:
                    output[card_type] = possible_hands

            if card_type == STRAIGHT:
                continue

        if any(len(suits) == 3 for num, suits in nums_dict.items()):
            card_type = THREE_OF_A_KIND
            output[card_type] += [{"cards": cards, "hand": hand}]

            if len(output[card_type]) > 1:
                max_hands = [max(num for num, suits in generate_num_dict(option["hand"]).items() if len(suits) == 3) for option in output[card_type]]
                max_of_max_hands = max(max_hands)
                possible_hands = [output[card_type][idx] for idx, max_hand in enumerate(max_hands) if max_hand == max_of_max_hands]

                if len(possible_hands) > 1:
                    max_hands = [sorted(card["num"] for card in option["cards"][-2:]) for option in possible_hands]
                    max_hands = [sorted(num for num, suits in generate_num_dict(option["hand"]).items() if len(suits) != 3) for option in output[card_type]]
                    if all(max_hand == max_hands[0] for max_hand in sorted(max_hands)):
                        max_hands = [max(card["num"] for card in option["cards"][-2:] if card["num"] != max_hands[0]) for option in possible_hands]
                    max_of_max_hands = max(max_hands)
                    output[card_type] = [possible_hands[idx] for idx, max_hand in enumerate(max_hands) if max_hand == max_of_max_hands]
                else:
                    output[card_type] = possible_hands

            if card_type == THREE_OF_A_KIND:
                continue

        if len([1 for num, suits in nums_dict.items() if len(suits) == 2]) == 2:
            card_type = TWO_PAIR
            output[card_type] += [{"cards": cards, "hand": hand}]

            if len(output[card_type]) > 1:
                card_pairs = [[num for num, suits in generate_num_dict(option["hand"]).items() if len(suits) == 2] for option in output[card_type]]
                highest_pair = max(card_pair[pair_num] for pair_num in range(2) for card_pair in card_pairs)
                possible_hands = [output[card_type][idx] for idx, card_pair in enumerate(card_pairs) if highest_pair in card_pair]

                if len(possible_hands) > 1:
                    for idx in range(5):
                        card_totals = [option["cards"][idx]["num"] for option in possible_hands]
                        if all(card_totals[0] == card_total for card_total in card_totals):
                            continue

                        max_hand = max(card_totals)
                        output[card_type] = [possible_hands[idx] for idx, card_total in enumerate(card_totals) if card_total == max_hand]
                        break
                else:
                    output[card_type] = possible_hands

            if card_type == TWO_PAIR:
                continue

        if len([1 for num, suits in nums_dict.items() if len(suits) == 2]) == 1:
            card_type = ONE_PAIR
            output[card_type] += [{"cards": cards, "hand": hand}]

            if len(output[card_type]) > 1:
                max_hands = [num for option in output[card_type] for num, suits in generate_num_dict(option["hand"]).items() if len(suits) == 2]
                max_of_max_hands = max(max_hands)
                possible_hands = [output[card_type][idx] for idx, max_hand in enumerate(max_hands) if max_hand == max_of_max_hands]

                if len(possible_hands) > 1:
                    max_hands = [max(num for num, suits in generate_num_dict(option["hand"]).items() if len(suits) != 2) for option in output[card_type]]
                    max_of_max_hands = max(max_hands)
                    output[card_type] = [possible_hands[idx] for idx, max_hand in enumerate(max_hands) if max_hand == max_of_max_hands]
                else:
                    output[card_type] = possible_hands

            if card_type == ONE_PAIR:
                continue

        if not card_type:
            card_type = HIGH_CARD
            output[card_type] += [{"cards": cards, "hand": hand}]

            if len(output[card_type]) > 1:
                cards_dict = defaultdict(list)
                temp_output = defaultdict(list[str])

                for option in output[card_type]:
                    for card in option["cards"]:
                        cards_dict[card["num"]].append(option)

                for card_num, options in cards_dict.items():
                    if len(options) > 1:
                        continue

                    temp_output[card_num] += options

                if temp_output:
                    output[card_type] = temp_output[max(temp_output.keys())]

            if card_type == HIGH_CARD:
                continue

    return output
