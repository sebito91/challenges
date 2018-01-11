#!/usr/bin/env python
import sys

def main():
    """ Main function to handle multiple chars in our test file.
    Make sure that we only output one char per repeat. """
    max_len = 250
    min_len = 1
    output = []
    last = ""

    for check in sys.stdin:
        if check and len(check.strip()) <= max_len and len(check) > min_len:
            for item in list(check.strip()):
                if item != last:
                    output.append(item)
                    last = item

    print "{}".format(''.join(output))

if __name__ == "__main__":
    main()
