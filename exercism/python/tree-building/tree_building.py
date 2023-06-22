"""Module to build a tree-structure for records."""

from collections import defaultdict


class Record:
    """The default object for a record."""

    def __init__(self, record_id: int, parent_id: int):
        """Initialize the Record object.

        :param record_id: int - ID for the particular record
        :param parent_id: int - ID for the new record's parent
        """
        self.record_id = record_id
        self.parent_id = parent_id


class Node:
    """The default object for a node."""

    def __init__(self, node_id: int):
        """Initialize the Node object.

        :param node_id: int - ID for the particular node.
        """
        self.node_id = node_id
        self.children = []


def BuildTree(records: list[object]) -> object:
    """Construct a tree-like object of Records and Nodes.

    :param records: list[object] - representation of the given record's layout within the tree, with ID + ParentID respectively
    :return: object - root node of the Record tree.
    """
    root = None
    if not records:
        return root

    records.sort(key=lambda x: x.record_id)
    ordered_id = [i.record_id for i in records]

    if (ordered_id[-1] != len(ordered_id) - 1) or (ordered_id[0] != 0):
        raise ValueError("Record id is invalid or out of order.")

    parents_and_children = defaultdict(list)
    for record in records:
        if record.record_id == 0 and record.parent_id != 0:
            raise ValueError("Node parent_id should be smaller than it's record_id.")

        if record.record_id < record.parent_id:
            raise ValueError("Node parent_id should be smaller than it's record_id.")

        if record.record_id == record.parent_id and record.record_id != 0:
            raise ValueError("Only root should have equal record and parent id.")

        parents_and_children[record.parent_id].append(Node(record.record_id))

    for parent, children in parents_and_children.items():
        for child in children:
            if child.node_id == parent:
                child.children = children[1:]
                root = child
                continue

            child.children = parents_and_children.get(child.node_id, [])

    return root
