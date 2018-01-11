#!/usr/bin/env python
import sys
from collections import deque

def main():
    """ Calculate the volume of fluid required for a 3D Printer """
    inputs = readfile()

    numpolys = int(inputs[0][0])
    polys = {}

    count = 1
    for poly in range(numpolys):
        facecount = int(inputs[count][0])
        polys[poly] = {'facecount': facecount,
                'faces': [inputs[count+a] for a in range(1,facecount+1)]
                }
        count += facecount + 1


    # iterate through the faces and find the absolute value distance between
    # probably wrong, but worth a shot
    volume = 0.0
    for each,value in polys.items():
        maxes = {'x': 0.0, 'y': 0.0, 'z': 0.0}
        for item in range(polys[each]['facecount']):
            x, y, z = calc_diff(polys[each]['faces'][item])

            maxes['x'] = max(maxes['x'], x)
            maxes['y'] = max(maxes['y'], y)
            maxes['z'] = max(maxes['z'], z)

        volume += maxes['x'] * maxes['y'] * maxes['z']

    print "{0:.2f}".format(volume)

def calc_diff(items):
    """ Receive list of list of points, determine the largest values """
    queue = deque(items)
    points = int(queue.popleft())
    thex = []
    they = []
    thez = []
    x = 0.0
    y = 0.0
    z = 0.0

    for each in range(1, points+1):
        thex.append(queue.popleft())
        they.append(queue.popleft())
        thez.append(queue.popleft())

    for each in range(points-1):
        x = max(abs(thex[each]-thex[each+1]), x)
        y = max(abs(they[each]-they[each+1]), y)
        z = max(abs(thez[each]-thez[each+1]), z)

    if points > 1:
        x = max(abs(thex[0]-thex[-1]), x)
        y = max(abs(they[0]-they[-1]), y)
        z = max(abs(thez[0]-thez[-1]), z)

    return x, y, z

def readfile():
    """ Function to read the files in from stdin """
    lines = []
    for line in sys.stdin:
        lines.append(line.split())

    for each in lines:
        for item in range(len(each)):
            each[item] = float(each[item])

    return lines

if __name__ == "__main__":
    main()
