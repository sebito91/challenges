"""Module to store high scores from the game Frogger."""


class HighScores:
    """Object to store high scores from the game Frogger."""

    def __init__(self, scores: list[int]) -> None:
        """Initialize the HighScores object.

        :param scores: list[int] - list of highest scores in Frogger.
        :return: NoneType
        """
        self._scores = scores

    def latest(self) -> int:
        """Return the latest high score to be added."""
        return self._scores[-1]

    def personal_best(self) -> int:
        """Return the highest personal best score."""
        return max(self._scores)

    def personal_top_three(self) -> int:
        """Return the personal best three scores."""
        return sorted(self._scores, reverse=True)[:3]

    @property
    def scores(self):
        """Return the set of high scores."""
        return self._scores

