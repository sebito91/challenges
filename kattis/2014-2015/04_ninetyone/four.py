#!/usr/bin/env python
import sys

def main():
    """ Call our recursive fun for McCarthy 91 """
    lines = sys.stdin.read().split()
    print lines

    for line in lines:
        print M(int(line.strip()))

def M(number):
    """ Recursive implementation of the McCarthy 91 """
    if number > 100:
        return number - 10

    return M(M(number + 11))

if __name__ == "__main__":
    main()
