"""Module to calculate the currency exchange given a variety of parameters and options."""


def exchange_money(budget: float, exchange_rate: float) -> float:
    """Calculate the value of the exchanged currency.

    :param budget: float - amount of money you are planning to exchange.
    :param exchange_rate: float - unit value of the foreign currency.
    :return: float - exchanged value of the foreign currency you can receive.

    Function to convert from one currency to another using the the given exchange rate.
    """

    return budget * (1 / exchange_rate)


def get_change(budget: float, exchanging_value: float) -> float:
    """Calculate the amount of money that is left from the budget.

    :param budget: float - amount of money you own.
    G:param exchanging_value: float - amount of your money you want to exchange now.
    :return: float - amount left of your starting currency after exchanging.

    Function to calculate the remaining amount of money from the exchange budget.
    """

    return budget - exchanging_value


def get_value_of_bills(denomination: int, number_of_bills: int) -> int:
    """Calculate the total value of bills that the booth would provide.

    :param denomination: int - the value of a bill.
    :param number_of_bills: int - amount of bills you received.
    :return: int - total value of bills you now have.

    Function to calculate the number of whole bills of given `denomination` that
    the booth would return to the customer.
    """

    return denomination * number_of_bills


def get_number_of_bills(budget: float, denomination: int) -> int:
    """Calculate the number of currency bills to receive given the budget and denomination.

    :param budget: float - the amount of money you are planning to exchange.
    :param denomination: int - the value of a single bill.
    :return: int - number of bills after exchanging all your money.

    Function to calculate the number of currency bills of `denomination` we can expect to
    receive given the specific `budget`.
    """

    return round(budget // denomination)


def get_leftover_of_bills(budget: float, denomination: int) -> float:
    """Calculate the remaining amount from the `budget` given the set of bills of `denomination`.

    :param budget: float - the amount of money you are planning to exchange.
    :param denomination: int - the value of a single bill.
    :return: float - the leftover amount that cannot be exchanged given the current denomination.

    Function to return the amount of the budget that is left over from splitting into bills
    of `denomination` value. This is the amount that goes to the house.
    """

    return budget - (get_number_of_bills(budget, denomination) * denomination)


def exchangeable_value(budget: float, exchange_rate: float, spread: int, denomination: int) -> int:
    """Calculate the maximum value of the new currency given the current spread.

    :param budget: float - the amount of your money you are planning to exchange.
    :param exchange_rate: float - the unit value of the foreign currency.
    :param spread: int - percentage that is taken as an exchange fee.
    :param denomination: int - the value of a single bill.
    :return: int - maximum value you can get.

    Function to calculate the entire exchanged amount when taking into account the `spread`.
    """

    return get_value_of_bills(denomination, get_number_of_bills(exchange_money(budget, exchange_rate * (1 + (spread / 100))), denomination))
