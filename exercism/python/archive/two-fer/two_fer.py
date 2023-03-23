"""Module to run the two-fer challenge."""


def two_fer(name: str = "you") -> str:
    """Process the given `name` into a phrase for the two-fer challange.

    :param name: str - provided name for the cookie
    :return: str - return the two-fer statement given the `name`

    Function to receive a name and provide a two-fer statement using that `name`.
    """

    return f"One for {name}, one for me."
