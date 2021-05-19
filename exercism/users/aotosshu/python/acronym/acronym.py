DELIMITERS = "-_"
TRANS = str.maketrans(DELIMITERS, " " * len(DELIMITERS))


def abbreviate(words: str) -> str:
    return "".join([word[0].upper()
                    for word
                    in words.translate(TRANS).split()])
