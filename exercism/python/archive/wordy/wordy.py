"""Module to process a written math question into an integer result using left-to-right processing."""

import operator


def answer(question: str) -> int:
    """Calculate the answer to the given written math question.

    :param question: str - basic math question to solve
    :return: int - value calculated using left-to-right processing instead of traditional math operations.
    """

    terms = {
        "multiplied": operator.mul,
        "plus": operator.add,
        "minus": operator.sub,
        "divided": operator.truediv,
    }

    # This solution will use relatively simple methods to address the problem.
    # Using tools like regex are a bit more involved for a simple challenge like this one.
    tokens = question.split()
    if not tokens[0] == "What" and tokens[1] == "is":
        raise ValueError("unknown operation")

    # start after the "What is" portion of the `question`
    result = None

    operation = None
    left_operand = None
    right_operand = None

    val = None
    last_token = None

    for idx, token in enumerate(tokens[2:]):
        if token in ["divided", "multiplied"]:
            if tokens[2 + idx + 1] != "by":
                raise ValueError("syntax error missing operator")

        if token == "by":
            continue

        if "?" in token:
            try:
                token = token.split("?")[0]
            except ValueError as exc:
                raise ValueError("syntax error") from exc

        try:
            # check if we are a value in our equation
            val = int(token)
        except ValueError as exc:
            # check if we've already assigned an op and bomb out if we did (e.g. What is 5 plus plus 5?)
            if operation:
                raise ValueError("syntax error") from exc

            operation = terms.get(token, None)
            if not operation:
                raise ValueError("unknown operation") from exc

            last_token = "operation"
        else:
            # if the previous token was a value, this is an error
            if last_token == "value":
                raise ValueError("syntax error")

            last_token = "value"
            if not left_operand:
                left_operand = val
            elif not right_operand:
                right_operand = val

        if left_operand and operation and right_operand and last_token == "value":
            result = operation(left_operand, right_operand)
            val = None
            operation = None
            left_operand = result
            right_operand = None

    # hack to handle malformed problems
    if not result and (not left_operand or operation):
        raise ValueError("syntax error")

    if not result and left_operand:
        return left_operand

    return result
