"""Module to transpone the set of rows into columns."""

from collections import defaultdict


def transpose(lines: list[str]) -> list[str]:
    """Transpose the given set of rows into columns.

    :param lines: list[str] - series of rows to convert to columns.
    :return: list[str] - series of columns from converted rows.
    """
    output = defaultdict(str)
    fill = " "

    words = lines.split("\n")
    max_len = max(words, key=len)
    lengths = [len(word) for word in words]

    for word_num, each_word in enumerate(words):
        word_len = len(each_word)
        for idx in range(len(max_len)):
            if word_len > idx:
                output[idx] += each_word[idx]
            elif any(idx < len_word for len_word in lengths[word_num:]):
                output[idx] += fill

    return "\n".join(output.values())
