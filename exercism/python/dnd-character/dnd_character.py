"""Module to generate characters for Dungeons 'n Dragons."""

import random


class Character:
    """Object to represent a DnD character."""

    def __init__(self):
        """Initialize the DnD Character object."""
        # generate the character's traits
        self._strength = self.ability()
        self._dexterity = self.ability()
        self._constitution = self.ability()
        self._intelligence = self.ability()
        self._wisdom = self.ability()
        self._charisma = self.ability()

        self._hitpoints = 10 + modifier(self.constitution)
        print(f"found myself: {self.__dict__}")

    @staticmethod
    def ability() -> int:
        """Generate a character's ability by rolling 4x 6-sided die and returning the sum of the top-three results.

        :return: int - sum of top-three results from rolling 4x 6-sided die
        """
        random.seed()
        return sum(sorted(random.randint(1, 6) for _ in range(4))[:3])

    @property
    def strength(self):
        """Return the strength of the character."""
        return self._strength

    @property
    def dexterity(self):
        """Return the dexterity of the character."""
        return self._dexterity

    @property
    def constitution(self):
        """Return the constitution of the character."""
        return self._constitution

    @property
    def intelligence(self):
        """Return the intelligence of the character."""
        return self._intelligence

    @property
    def wisdom(self):
        """Return the wisdom of the character."""
        return self._wisdom

    @property
    def charisma(self):
        """Return the charisma of the character."""
        return self._charisma

    @property
    def hitpoints(self):
        """Return the hitpoints of the character."""
        return self._hitpoints


def modifier(constitution: int) -> int:
    """Modify the character's constitution and return the result.

    :param constitution: int - the character's constitution
    :return: int - modified constitution count.
    """
    return (constitution - 10) // 2
