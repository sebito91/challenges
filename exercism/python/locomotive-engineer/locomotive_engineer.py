"""Functions which helps the locomotive engineer to keep track of the train."""


def get_list_of_wagons(*wagons: int) -> list[int]:
    """Return a list of wagons.

    :param: arbitrary number of wagons.
    :return: list - list of wagons.
    """

    wagon_list = []
    for wagon in wagons:
        wagon_list.append(wagon)

    return wagon_list


def fix_list_of_wagons(each_wagons_id: list[int], missing_wagons: list[int]) -> list[int]:
    """Fix the list of wagons.

    :parm each_wagons_id: list - the list of wagons.
    :parm missing_wagons: list - the list of missing wagons.
    :return: list - list of wagons.
    """

    incorrect_wagon_one, incorrect_wagon_two = each_wagons_id[0], each_wagons_id[1]
    return_wagon_list = [1] + missing_wagons

    for idx in range(3, len(each_wagons_id)):
        return_wagon_list.append(each_wagons_id[idx])

    return_wagon_list.append(incorrect_wagon_one)
    return_wagon_list.append(incorrect_wagon_two)

    return return_wagon_list


def add_missing_stops(route: dict, **stops) -> dict:
    """Add missing stops to route dict.

    :param route: dict - the dict of routing information.
    :param stops: arbitrary number of stops.
    :return: dict - updated route dictionary.
    """

    route["stops"] = list(stops.values())
    return route


def extend_route_information(route: dict, more_route_information: dict) -> dict:
    """Extend route information with more_route_information.

    :param route: dict - the route information.
    :param more_route_information: dict -  extra route information.
    :return: dict - extended route information.
    """

    for key, value in more_route_information.items():
        route[key] = value

    return route


def fix_wagon_depot(wagons_rows: list[list[tuple]]) -> list[list[tuple]]:
    """Fix the list of rows of wagons.

    :param wagons_rows: list[list[tuple]] - the list of rows of wagons.
    :return: list[list[tuple]] - list of rows of wagons.
    """

    wagons_output = []
    for wagons_row in wagons_rows:
        for col_num, wagon in enumerate(wagons_row):
            try:
                _ = wagons_output[col_num]
            except IndexError:
                wagons_output.append([])

            wagons_output[col_num].append(wagon)

    return wagons_output
