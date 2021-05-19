"""A Program to determine whether a given number is valid per the Luhn Formula"""


class Luhn:
    def __init__(self, card_num):
        self.card_num = card_num.replace(' ', '')
        self.first_step = self.card_num[-2::-2]
        self.second_half = self.card_num[-1::-2]

    def valid(self):
        """Validate the given number per Luhn formula."""
        if not self.card_num.isdigit():
            return False
        luhn_sum = 0
        for i in self.first_step:
            if int(i) * 2 > 9:
                luhn_sum += int(i) * 2 - 9
            else:
                luhn_sum += int(i) * 2
        for b in self.second_half:
            luhn_sum += int(b)
        if luhn_sum == 0 or luhn_sum % 10 == 0:
            return True
        else:
            return False
