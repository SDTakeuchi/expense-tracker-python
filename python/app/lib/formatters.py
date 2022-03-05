class Formatter:
    @staticmethod
    def truncate(text, max_len, trailing_char='...'):
        if not isinstance(text, str):
            raise Exception(
                f'Formatter.truncate() received not-str first arg; type:{type(text)}'
            )
        if not isinstance(max_len, int):
            raise Exception(
                f'Formatter.truncate() received not-int second arg; type:{type(max_len)}'
            )
        if text > max_len:
            return text[:max_len] + trailing_char
        return text

    @staticmethod
    def half_round_num(number, hide_extra_zero=False):
        """
        round numbers to nearest 0.5
        """
        num_rounded = round(number * 2) / 2
        result_num = int(num_rounded) if hide_extra_zero is True and num_rounded % 1 == 0 else num_rounded
        return result_num