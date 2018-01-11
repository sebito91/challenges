#!/usr/bin/env python
import sys

def main():
    """ Determine the final score between two players in a round-based (trick)
    card game. """
    counter = 0
    inputs = ['','','']
    score_one = 0
    score_two = 0

    for line in sys.stdin:
        # Line 1 is the number of cards to expect for each player (1 <= N <= 1000)
        if counter == 0:
            inputs[counter] = line.strip()
        else:
            inputs[counter] = list(line.strip().replace(' ', ''))
            for each in range(int(inputs[0])):
                inputs[counter][each] = inputs[counter][each].replace('A', '13')
                inputs[counter][each] = inputs[counter][each].replace('K', '12')
                inputs[counter][each] = inputs[counter][each].replace('Q', '11')
                inputs[counter][each] = inputs[counter][each].replace('J', '10')

        counter += 1

    for card in range(int(inputs[0])):
        # if they're the same, do nothing
        if int(inputs[1][card]) == int(inputs[2][card]):
            continue

        # if A is greater, plus one:
        if int(inputs[1][card]) > int(inputs[2][card]):
            score_one += 1
        else:
            score_two += 1

    if score_one > score_two:
        print "PLAYER 1 WINS"
    elif score_two > score_one:
        print "PLAYER 2 WINS"
    else:
        print "TIE"

if __name__ == "__main__":
    main()
