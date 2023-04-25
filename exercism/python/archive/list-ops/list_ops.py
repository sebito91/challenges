"""Module to implement typical list functions without using standard libary operations."""


from typing import Callable


def append(list1: list, list2: list) -> list:
    """Append the items from `list2` onto `list1` (duplicates allowed).

    :param list1: list - primary list for input
    :param list2: list - list of items to append to `list1`
    :return: list - extended list of items between `list1` and `list2`
    """
    return list1 + list2


def concat(lists: list) -> list:
    """Concatenate the series of lists into one longer list.

    :param lists: list - series of lists to concatenate
    :return: list - concatenated list of elements
    """
    output = []
    for list_of_items in lists:
        output += list_of_items

    return output


def filter(function: Callable, list_of_items: list) -> list:
    """Run the given `list_of_items` through the provided `function` and return the series of elements that are True.

    :param function: Callable - predicate to run the `list_of_items` through.
    :param list_of_items: list - list of items to run through `function`.
    :return: list - series of values where the `function` evaluates to True.
    """
    return [item for item in list_of_items if function(item)]


def length(list_of_items: list) -> int:
    """Return the total number of items in the provided `list_of_items`.

    :param list_of_items: list - series of items to count.
    :return: int - length of provided `list_of_items`.
    """
    return sum(1 for item in list_of_items)


def map(function: Callable, list_of_items: list) -> list:
    """Return the list of results when running `list_of_items` through the given `function`.

    :param function: Callable - the function to map each of the `list_of_items` through.
    :param list_of_items: list - series of values to run through the given `function`.
    :return: list - results of running `list_of_items` through the given `function`.
    """
    return [function(item) for item in list_of_items]


def foldl(function: Callable, list_of_items: list, initial: int) -> int:
    """Reduce each item into the accumulator from the left.

    :param function: Callable - accumulator function.
    :param list_of_items: list - list of items to run through the accumulator.
    :param initial: int - initial value to run through the accumulator.
    :return: int - reduced value of the accumulator.
    """
    val = initial
    for item in list_of_items:
        val = function(val, item)

    return val


def foldr(function: Callable, list_of_items: list, initial: int) -> int:
    """Reduce each item into the accumulator from the right.

    :param function: Callable - accumulator function.
    :param list_of_items: list - list of items to run through the accumulator.
    :param initial: int - initial value to run through the accumulator.
    :return: int - reduced value of the accumulator.
    """
    val = initial
    for item in list_of_items[::-1]:
        val = function(item, val)

    return val


def reverse(list_of_items: list) -> list:
    """Return the provided list in reversed order.

    :param list_of_items: list - list of items to reverse
    :return: list - reversed list of items.
    """
    return list_of_items[::-1]
