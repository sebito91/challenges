"""Module to use binary-search to find the title of a song."""


def find(search_list: list[int], value: int) -> int:
    """Find the given value in a set of search items.

    :param search_list: list[int] - list of integers to represent song titles.
    :param value: int - specific song title to search for, in `int` form.
    :return: int - position in original `search_list` where item is found.
    """

    if not search_list or value < search_list[0] or value > search_list[-1]:
        raise ValueError("value not in array")

    mid = len(search_list) // 2
    if search_list[mid] == value:
        return mid

    if value < search_list[mid]:
        return find(search_list[:mid], value)

    return (mid + 1) + find(search_list[mid + 1:], value)
