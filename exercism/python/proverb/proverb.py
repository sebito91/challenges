"""Module to pull together the horseshoe proverb."""


def proverb(*elements: list[str], qualifier: str = None) -> list[str]:
    """Generate the horseshoe proverb.

    :param elements: list[str] - series of elements for the proverb.
    :param qualifier: str - optional qualifier for the proverb.
    :return: list[str] - generated set of proverb lines.
    """
    output = []
    final_line = "And all for the want of a"
    if qualifier:
        final_line += f" {qualifier}"

    for idx in range(len(elements) - 1):
        output.append(f"For want of a {elements[idx]} the {elements[idx + 1]} was lost.")

    if elements:
        output.append(f"{final_line} {elements[0]}.")

    return output
