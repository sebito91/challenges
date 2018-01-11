#!/usr/bin/env python
import sys

def main():
    """ Handle the 2048 game based off of inputs!
        Move in this case can be one of four directions:
        0 - left
        1 - up
        2 - right
        3 - down
    """
    inputs = readfile()
    outputs = []
    move = int(inputs[4][0])

    if move == 0:
        outputs = move_left(inputs)
    elif move == 1:
        outputs = move_up(inputs)
    elif move == 2:
        outputs = move_right(inputs)
    else:
        outputs = move_down(inputs)

    for item in range(4):
        output = ""
        for each in outputs[item]:
            output += str(each) + " "
        print output

def move_down(inputs):
    """ Make our moves based on the direction given
        Transpose, then shift left
    """
    worklist = []
    for i in range(4):
        new = []
        for j in range(4):
            new.append(inputs[3-j][i])
        worklist.append(new)

    inputs = worklist

    # iterate rows
    for item in range(4):
        workrow = inputs[item]
        #print workrow

        counter = 0
        while workrow[0] == 0 and counter < 4:
            workrow[0] = workrow[counter]
            workrow[counter] = 0
            counter += 1

        if workrow[0] == workrow[1]:
            workrow[0] += workrow[1]
            workrow[1] = workrow[2]
            workrow[2] = workrow[3]
            workrow[3] = 0
        elif workrow[0] == workrow[2] and workrow[1] == 0:
            workrow[0] += workrow[2]
            workrow[1] = workrow[3]
            workrow[2] = 0
            workrow[3] = 0
        elif workrow[0] == workrow[3] and workrow[1] == 0 and workrow[2] == 0:
            workrow[0] += workrow[3]
            workrow[1] = 0
            workrow[2] = 0
            workrow[3] = 0

        if workrow[1] == 0:
            workrow[1] = workrow[2]
            workrow[2] = workrow[3]
            workrow[3] = 0

        if workrow[1] == workrow[2]:
            workrow[1] += workrow[2]
            workrow[2] = workrow[3]
            workrow[3] = 0
        elif workrow[1] == workrow[3] and workrow[2] == 0:
            workrow[1] += workrow[3]
            workrow[2] = 0
            workrow[3] = 0

        if workrow[2] == 0:
            workrow[2] = workrow[3]
            workrow[3] = 0
        elif workrow[2] == workrow[3]:
            workrow[2] += workrow[3]
            workrow[3] = 0

        inputs[item] = workrow

    worklist = []
    for i in range(4):
        new = []
        for j in range(4):
            new.append(inputs[j][3-i])
        worklist.append(new)

    inputs = worklist

    return inputs[0:4]

def move_up(inputs):
    """ Make our moves based on the direction given
        Transpose, then shift right
    """
    worklist = []
    for i in range(4):
        new = []
        for j in range(4):
            new.append(inputs[3-j][i])
        worklist.append(new)

    inputs = worklist

    # iterate columns
    for item in range(4):
        workrow = inputs[item]
        #print workrow

        counter = 2
        while workrow[3] == 0 and counter >= 0:
            workrow[3] = workrow[counter]
            workrow[counter] = 0
            counter -= 1

        if workrow[3] == workrow[2]:
            workrow[3] += workrow[2]
            workrow[2] = workrow[1]
            workrow[1] = workrow[0]
            workrow[0] = 0
        elif workrow[3] == workrow[1] and workrow[2] == 0:
            workrow[3] += workrow[1]
            workrow[2] = workrow[0]
            workrow[1] = 0
            workrow[0] = 0
        elif workrow[3] == workrow[0] and workrow[1] == 0 and workrow[2] == 0:
            workrow[3] += workrow[0]
            workrow[2] = 0
            workrow[1] = 0
            workrow[0] = 0

        if workrow[2] == 0:
            workrow[2] = workrow[1]
            workrow[1] = workrow[0]
            workrow[0] = 0

        if workrow[2] == workrow[1]:
            workrow[2] += workrow[1]
            workrow[1] = workrow[0]
            workrow[0] = 0
        elif workrow[2] == workrow[0] and workrow[1] == 0:
            workrow[2] += workrow[0]
            workrow[1] = 0
            workrow[0] = 0

        if workrow[1] == 0:
            workrow[1] = workrow[0]
            workrow[0] = 0
        elif workrow[1] == workrow[0]:
            workrow[1] += workrow[0]
            workrow[0] = 0

        inputs[item] = workrow

    worklist = []
    for i in range(4):
        new = []
        for j in range(4):
            new.append(inputs[j][3-i])
        worklist.append(new)

    inputs = worklist

    return inputs[0:4]

def move_right(inputs):
    """ Make our moves based on the direction given """

    # iterate rows
    for item in range(4):
        workrow = inputs[item]
        #print workrow

        counter = 2
        while workrow[3] == 0 and counter >= 0:
            workrow[3] = workrow[counter]
            workrow[counter] = 0
            counter -= 1

        if workrow[3] == workrow[2]:
            workrow[3] += workrow[2]
            workrow[2] = workrow[1]
            workrow[1] = workrow[0]
            workrow[0] = 0
        elif workrow[3] == workrow[1] and workrow[2] == 0:
            workrow[3] += workrow[1]
            workrow[2] = workrow[0]
            workrow[1] = 0
            workrow[0] = 0
        elif workrow[3] == workrow[0] and workrow[1] == 0 and workrow[2] == 0:
            workrow[3] += workrow[0]
            workrow[2] = 0
            workrow[1] = 0
            workrow[0] = 0

        if workrow[2] == 0:
            workrow[2] = workrow[1]
            workrow[1] = workrow[0]
            workrow[0] = 0

        if workrow[2] == workrow[1]:
            workrow[2] += workrow[1]
            workrow[1] = workrow[0]
            workrow[0] = 0
        elif workrow[2] == workrow[0] and workrow[1] == 0:
            workrow[2] += workrow[0]
            workrow[1] = 0
            workrow[0] = 0

        if workrow[1] == 0:
            workrow[1] = workrow[0]
            workrow[0] = 0
        elif workrow[1] == workrow[0]:
            workrow[1] += workrow[0]
            workrow[0] = 0

        inputs[item] = workrow

    return inputs[0:4]

def move_left(inputs):
    """ Make our moves based on the direction given """

    # iterate rows
    for item in range(4):
        workrow = inputs[item]
        #print workrow

        counter = 0
        while workrow[0] == 0 and counter < 4:
            workrow[0] = workrow[counter]
            workrow[counter] = 0
            counter += 1

        if workrow[0] == workrow[1]:
            workrow[0] += workrow[1]
            workrow[1] = workrow[2]
            workrow[2] = workrow[3]
            workrow[3] = 0
        elif workrow[0] == workrow[2] and workrow[1] == 0:
            workrow[0] += workrow[2]
            workrow[1] = workrow[3]
            workrow[2] = 0
            workrow[3] = 0
        elif workrow[0] == workrow[3] and workrow[1] == 0 and workrow[2] == 0:
            workrow[0] += workrow[3]
            workrow[1] = 0
            workrow[2] = 0
            workrow[3] = 0

        if workrow[1] == 0:
            workrow[1] = workrow[2]
            workrow[2] = workrow[3]
            workrow[3] = 0

        if workrow[1] == workrow[2]:
            workrow[1] += workrow[2]
            workrow[2] = workrow[3]
            workrow[3] = 0
        elif workrow[1] == workrow[3] and workrow[2] == 0:
            workrow[1] += workrow[3]
            workrow[2] = 0
            workrow[3] = 0

        if workrow[2] == 0:
            workrow[2] = workrow[3]
            workrow[3] = 0
        elif workrow[2] == workrow[3]:
            workrow[2] += workrow[3]
            workrow[3] = 0

        inputs[item] = workrow

    return inputs[0:4]

def check_digits(source, dest, num_left):
    """ Irregardless of the direction, what's the difference between these """
    if source == 0:
        source = dest
        dest = 0
    elif source == dest:
        source += dest
        dest = 0

    return source, dest

def readfile():
    """ Handler to read the inputs from the source file,
    return None if there's a problem """
    lines = []

    for line in sys.stdin:
        lines.append(line.split())

    for each in lines:
        for item in range(len(each)):
            each[item] = int(each[item])

    if not lines:
        return None

    return lines

if __name__ == "__main__":
    main()
