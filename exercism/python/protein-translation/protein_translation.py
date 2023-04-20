"""Module to translate RNA codons to amino acids."""


PROTEINS = {
    "AUG": "Methionine",
    "UUU": "Phenylalanine",
    "UUC": "Phenylalanine",
    "UUA": "Leucine",
    "UUG": "Leucine",
    "UCU": "Serine",
    "UCC": "Serine",
    "UCA": "Serine",
    "UCG": "Serine",
    "UAU": "Tyrosine",
    "UAC": "Tyrosine",
    "UGU": "Cysteine",
    "UGC": "Cysteine",
    "UGG": "Tryptophan",
    "UAA": "STOP",
    "UAG": "STOP",
    "UGA": "STOP",
}


def proteins(strand: str) -> list[str]:
    """Translate the given `strand` of RNA codons to their formal amino acid names.

    :param strand: str - strans of RNA codons to translate
    :return: list[str] - parsed list of amino acids
    """

    output = []
    for idx in range(0, len(strand), 3):
        amino = PROTEINS.get(strand[idx:idx + 3])
        if amino == "STOP":
            break

        output.append(amino)

    return output
