"""Module to implement a simple linked list."""


class Node:
    """Node object for individual components of the linked list."""

    def __init__(self, value: int):
        """Initialize the particular node.

        :param value: int - value to store in the data
        """
        self._data = value
        self._next = None

    def value(self) -> int:
        """Return the data value of the node.

        :return: int - value stored within the node
        """
        return self._data

    def next(self) -> object:
        """Return the next node in the list.

        :return: Node - pointer to the next Node in the list
        """
        return self._next


class LinkedListIterator:
    """Object to help iterate through a LinkedList object."""

    def __init__(self, linked_list: object):
        """Initialize the LinkedList iterator object given the linked_list.

        :param linked_list: object - the LinkedList object to create an iterator with
        """
        self.current_node = linked_list._head

    def __iter__(self) -> object:
        """Instantiate the iterator for the LinkedList."""
        return self

    def __next__(self) -> object:
        """Return the next object in the LinkedList."""
        if not self.current_node:
            raise StopIteration

        current_node_value = self.current_node.value()
        self.current_node = self.current_node.next()
        return current_node_value


class LinkedList:
    """Object to hold a linked list."""

    def __init__(self, values: list[int] = []):
        """Initialize the linked list given the specific values.

        :param values: list[int] - list of integers to represent song IDs
        """
        self._len = 0
        self._head = None

        for value in values:
            self.push(value)

    def __len__(self) -> int:
        """Return the length of the LinkedList."""
        return self._len

    def __iter__(self) -> object:
        """Iterate the items of the LinkedList."""
        return LinkedListIterator(self)

    def head(self) -> object:
        """Return the head of the LinkedList."""
        if not self._head:
            raise EmptyListException("The list is empty.")
        return self._head

    def push(self, value: int):
        """Push the new value onto the LinkedList.

        :param value: int - new song ID for the list
        """
        new_node = Node(value)
        new_node._next = self._head
        self._head = new_node
        self._len += 1
        print(f"value: {value}, head: {self._head}, len: {self._len}, new_node: {new_node}")

    def pop(self) -> object:
        """Remove the last element of the LinkedList."""
        if len(self) <= 0:
            raise EmptyListException("The list is empty.")

        pop_node = self._head
        self._head = self._head.next()
        self._len -= 1

        return pop_node.value()

    def reversed(self) -> object:
        """Reverse the LinkedList."""
        return list(self)[::-1]


class EmptyListException(Exception):
    """Custom Exception object for the LinkedList."""

    def __init__(self, message: str):
        """Initialize the EmptyListException object with the provided message.

        :param message: str - message to add into the custom exception
        """
        self.message = message
