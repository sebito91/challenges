"""Functions to keep track and alter inventory."""


def create_inventory(items: list[str]) -> dict:
    """Create a dict that tracks the amount (count) of each element on the `items` list.

    :param items: list - list of items to create an inventory from.
    :return: dict - the inventory dictionary.
    """

    inventory = {}
    for item in items:
        if inventory.get(item, None):
            inventory[item] += 1
        else:
            inventory[item] = 1

    return inventory


def add_items(inventory: dict, items: list[str]) -> dict:
    """Add or increment items in inventory using elements from the items `list`.

    :param inventory: dict - dictionary of existing inventory.
    :param items: list - list of items to update the inventory with.
    :return: dict - the inventory updated with the new items.
    """

    new_items = create_inventory(items)
    for item, count in inventory.items():
        if new_items.get(item):
            new_items[item] += count
        else:
            new_items[item] = count

    return new_items


def decrement_items(inventory: dict, items: list[str]) -> dict:
    """Decrement items in inventory using elements from the `items` list.

    :param inventory: dict - inventory dictionary.
    :param items: list - list of items to decrement from the inventory.
    :return: dict - updated inventory with items decremented.
    """

    for item in items:
        if inventory.get(item) and inventory[item] > 0:
            inventory[item] -= 1

    return inventory


def remove_item(inventory: dict, item: str):
    """Remove item from inventory if it matches `item` string.

    :param inventory: dict - inventory dictionary.
    :param item: str - item to remove from the inventory.
    :return: dict - updated inventory with item removed. Current inventory if item does not match.
    """

    _ = inventory.pop(item, None)
    return inventory


def list_inventory(inventory: dict) -> list[tuple]:
    """Create a list containing all (item_name, item_count) pairs in inventory.

    :param inventory: dict - an inventory dictionary.
    :return: list of tuples - list of key, value pairs from the inventory dictionary.
    """

    return [(key, val) for key, val in inventory.items() if val > 0]
