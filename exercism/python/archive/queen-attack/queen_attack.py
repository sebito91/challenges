"""Module to confirm whether two Queens in Chess can attack one another."""


class Queen:
    """Class to check whether Queens can attack one another."""

    def __init__(self, row: int, column: int) -> object:
        """Initialize the object.

        :param row: int - row on the board where the queen lives.
        :param column: int - column on the board where the queen lives.
        :return: Queen - return an instance of the class.
        """
        self._check_row_col(row, column)

        self._row = row
        self._column = column

    @property
    def row(self):
        """Return the row for this Queen."""
        return self._row

    @property
    def column(self):
        """Return the column for this Queen."""
        return self._column

    @property
    def position(self):
        """Return the position of the Queen."""
        return (self.row, self.column)

    @staticmethod
    def _check_row_col(row: int, column: int):
        """Check whether the row and column are on the board.

        :param row: int - row for the chess board.
        :param column: int - column for the chess board.
        """
        if column > 7:
            raise ValueError("column not on board")

        if column < 0:
            raise ValueError("column not positive")

        if row < 0:
            raise ValueError("row not positive")

        if row > 7:
            raise ValueError("row not on board")

    def can_attack(self, another_queen: object) -> bool:
        """Check whether the given Queen can attack this Queen object.

        :param another_queen: Queen - new Queen to check whether we can attack.
        :return: bool - whether we can attack the new Queen in `another_queen`.
        """
        if self.position == another_queen.position:
            raise ValueError("Invalid queen position: both queens in the same square")

        if self.row == another_queen.row or self.column == another_queen.column:
            return True

        # diagonal set of options
        options = set()
        for col in range(8):
            row_up = self.row - (col + 1)
            row_down = self.row + (col + 1)
            col_up = self.column - (col + 1)
            col_down = self.column + (col + 1)

            if row_up >= 0 and col_up >= 0:
                options.add((row_up, col_up))

            if row_down < 8 and col_up >= 0:
                options.add((row_down, col_up))

            if row_up >= 0 and col_down < 8:
                options.add((row_up, col_down))

            if row_down < 8 and col_down < 8:
                options.add((row_down, col_down))

        if {another_queen.position} & options:
            return True

        return False
