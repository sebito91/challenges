"""Module to print out a given letter in diamond shape."""


def rows(letter: str) -> list[str]:
    """Print out the letters up to the given letter in a nice diamond shape.

    :param letter: str - input letter to print out in diamond form
    :return: list[str] - diamond shape of output up to `letter`
    """
    max_ord = ord(letter)
    distance = max_ord - ord("A")
    fill = " "

    output = []
    for x in range(distance, -1, -1):
        if distance - x == 0:
            output.append(f"{fill * x}{chr(ord('A') + (distance - x))}{fill * x}")
            continue

        output.append(f"{fill * x}{chr(ord('A') + (distance - x))}{fill * (2 * (distance - x) - 1)}{chr(ord('A') + (distance - x))}{fill * x}")

    if distance > 0:
        output += output[len(output) - 2::-1]

    return output
