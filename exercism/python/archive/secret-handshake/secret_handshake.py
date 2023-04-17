"""Module to test out the secret handshake encoding."""


def commands(binary_str: str) -> list[str]:
    """Calculate the secret handshake instructions and return the decoded list.

    :param binary_str: str - input string of binary "digit" to decode
    :return: list[str] - action sequence as decoded from the digit
    """

    actions = {
        0: "wink",
        1: "double blink",
        2: "close your eyes",
        3: "jump",
        4: "reverse",
    }

    output = []
    for idx, elem in enumerate(binary_str[::-1]):
        if elem == "1" and idx < 4:
            output.append(actions.get(idx, ""))

        if elem == "1" and idx == 4:
            output = output[::-1]

    return output
