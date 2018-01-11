#!/usr/bin/env python
import sys

def main():
    """ Output the number of days < 0C """
    num_temps = 0
    counter = 0
    inputs = ['','']
    cold_days = 0

    for line in sys.stdin:
        inputs[counter] = line.strip()
        counter += 1

    inputs[1] = inputs[1].split()
    for temps in range(int(inputs[0])):
        if int(inputs[1][temps]) < 0:
            cold_days += 1

    print cold_days

if __name__ == "__main__":
    main()
