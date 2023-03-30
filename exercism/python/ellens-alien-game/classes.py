"""Solution to Ellen's Alien Game exercise."""


class Alien:
    """Create an Alien object with location x_coordinate and y_coordinate.

    Attributes
    ----------
    (class)total_aliens_created: int
    x_coordinate: int - Position on the x-axis.
    y_coordinate: int - Position on the y-axis.
    health: int - Amount of health points.

    Methods
    -------
    hit(): Decrement Alien health by one point.
    is_alive(): Return a boolean for if Alien is alive (if health is > 0).
    teleport(new_x_coordinate, new_y_coordinate): Move Alien object to new coordinates.
    collision_detection(other): Implementation TBD.
    """

    total_aliens_created = 0

    def __init__(self, x_coordinate: int, y_coordinate: int):
        """Instantiate an object of Alien class."""
        self.x_coordinate = x_coordinate
        self.y_coordinate = y_coordinate
        self.health = 3

        Alien.total_aliens_created += 1

    def hit(self):
        """Decrement the health of the Alien instance by one."""
        if self.health > 0:
            self.health -= 1

    def is_alive(self) -> bool:
        """Check if the health of the Alien object is alive."""
        return self.health > 0

    def teleport(self, x_coordinate: int, y_coordinate: int):
        """Update the position for the Alien object."""
        self.x_coordinate = x_coordinate
        self.y_coordinate = y_coordinate

    def collision_detection(self, other_alien: object):
        """Check if the other object collides with this Alien object instance."""
        return


def new_aliens_collection(alien_start_positions: list[tuple]) -> list[Alien]:
    """Create a list of aliens of type Alien and return."""
    aliens = []
    for position in alien_start_positions:
        aliens.append(Alien(position[0], position[1]))

    return aliens
