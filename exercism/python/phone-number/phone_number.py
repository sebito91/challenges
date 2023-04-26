"""Module to validate a given phone number for NANP compliance."""


class PhoneNumber:
    """Class to validate a given Phone Number against NANP rules."""

    def __init__(self, number: str) -> object:
        """Initialize the PhoneNumber class object.

        :param number: str - NANP phone number string to validate.
        :return: PhoneNumber - object
        """
        clean_number = []
        for digit in number:
            try:
                _ = int(digit)
            except ValueError as exc:

                if digit.isalpha():
                    raise ValueError("letters not permitted") from exc

                if digit.isspace() or digit in ["(", ")", "-", "+", "."]:
                    continue

                raise ValueError("punctuations not permitted") from exc

            clean_number.append(digit)

        print(f"number: {number}, clean_number: {clean_number}")
        if len(clean_number) < 10:
            raise ValueError("must not be fewer than 10 digits")

        if len(clean_number) > 11:
            raise ValueError("must not be greater than 11 digits")

        if len(clean_number) == 11 and clean_number[0] != "1":
            raise ValueError("11 digits must start with 1")

        if len(clean_number) == 11:
            clean_number = clean_number[1:]

        if clean_number[0] == "0":
            raise ValueError("area code cannot start with zero")

        if clean_number[0] == "1":
            raise ValueError("area code cannot start with one")

        if clean_number[3] == "0":
            raise ValueError("exchange code cannot start with zero")

        if clean_number[3] == "1":
            raise ValueError("exchange code cannot start with one")

        self._number = ''.join(clean_number)
        self._area_code = ''.join(clean_number[0:3])
        self.exchange = ''.join(clean_number[3:6])
        self.remainder = ''.join(clean_number[6:])

    @property
    def number(self) -> str:
        """Return the stored number.

        :return: str - NANP phone number.
        """
        return self._number

    @property
    def area_code(self) -> str:
        """Return the stored area code.

        :return: str - NANP phone area code.
        """
        return self._area_code

    def pretty(self) -> str:
        """Generate a pretty version of the NANP phone number.

        :return: str - pretty version of number.
        """
        return f"({self.area_code})-{self.exchange}-{self.remainder}"
