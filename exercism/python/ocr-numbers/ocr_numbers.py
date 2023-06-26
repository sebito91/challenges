"""Module to convert the given OCR number to decimal."""


MODELS = {
    " _ | ||_|   ": "0",
    "     |  |   ": "1",
    " _  _||_    ": "2",
    " _  _| _|   ": "3",
    "   |_|  |   ": "4",
    " _ |_  _|   ": "5",
    " _ |_ |_|   ": "6",
    " _   |  |   ": "7",
    " _ |_||_|   ": "8",
    " _ |_| _|   ": "9",
}

LEN_CHAR = 3
LEN_LINE = 4


def convert(input_grid: list[str]) -> str:
    """Convert the given input_grid to the decimal value or return whether the number is garbled.

    :param input_grid: list[str] - OCR number to check
    :return: str - the converted OCR number to decimal, or a "?" if garbled
    """
    if len(input_grid) % 4 != 0:
        raise ValueError("Number of input lines is not a multiple of four")

    if not all(len(line) % 3 == 0 for line in input_grid):
        raise ValueError("Number of input columns is not a multiple of three")

    if not all(len(input_grid[0]) == len(line) for line in input_grid):
        raise ValueError("Number of input columns should match across all rows")

    output = []
    num_chars = len(input_grid[0]) // LEN_CHAR
    num_nums = len(input_grid) // LEN_LINE

    for num_row in range(num_nums):
        inner_output = []

        for num_char in range(num_chars):
            try:
                line = "".join(chars[(num_char * LEN_CHAR):(num_char * LEN_CHAR) + LEN_CHAR] for chars in input_grid[num_row * LEN_LINE:num_row * LEN_LINE + LEN_LINE])
                inner_output.append(MODELS[line])
            except KeyError:
                inner_output.append("?")

        output.append("".join(inner_output))

    return ",".join(output)
