"""Functions to manage and organize queues at Chaitana's roller coaster."""


EXPRESS_QUEUE = 1
NORMAL_QUEUE = 0


def add_me_to_the_queue(express_queue: list, normal_queue: list, ticket_type: int, person_name: str) -> list:
    """Add a person to the 'express' or 'normal' queue depending on the ticket number.

    :param express_queue: list - names in the Fast-track queue.
    :param normal_queue: list - names in the normal queue.
    :param ticket_type: int - type of ticket. 1 = express, 0 = normal.
    :param person_name: str - name of person to add to a queue.
    :return: list - the (updated) queue the name was added to.
    """

    if ticket_type == EXPRESS_QUEUE:
        return express_queue + [person_name]

    return normal_queue + [person_name]


def find_my_friend(queue: list, friend_name: str) -> int:
    """Search the queue for a name and return their queue position (index).

    :param queue: list - names in the queue.
    :param friend_name: str - name of friend to find.
    :return: int - index at which the friends name was found.
    """

    try:
        return queue.index(friend_name)
    except ValueError:
        return -1


def add_me_with_my_friends(queue: list, index: int, person_name: str) -> list:
    """Insert the late arrival's name at a specific index of the queue.

    :param queue: list - names in the queue.
    :param index: int - the index at which to add the new name.
    :param person_name: str - the name to add.
    :return: list - queue updated with new name.
    """

    if index > len(queue):
        return queue

    return queue[:index] + [person_name] + queue[index:]


def remove_the_mean_person(queue: list, person_name: str) -> list:
    """Remove the mean person from the queue by the provided name.

    :param queue: list - names in the queue.
    :param person_name: str - name of mean person.
    :return: list - queue update with the mean persons name removed.
    """

    if person_name not in queue:
        return queue

    idx = queue.index(person_name)
    return queue[:idx] + queue[idx + 1:]


def how_many_namefellows(queue: list, person_name: str) -> int:
    """Count how many times the provided name appears in the queue.

    :param queue: list - names in the queue.
    :param person_name: str - name you wish to count or track.
    :return: int - the number of times the name appears in the queue.
    """

    return queue.count(person_name)


def remove_the_last_person(queue: list) -> str:
    """Remove the person in the last index from the queue and return their name.

    :param queue: list - names in the queue.
    :return: str - name that has been removed from the end of the queue.
    """

    return queue.pop(-1)


def sorted_names(queue: list) -> list:
    """Sort the names in the queue in alphabetical order and return the result.

    :param queue: list - names in the queue.
    :return: list - copy of the queue in alphabetical order.
    """

    return sorted(queue)
