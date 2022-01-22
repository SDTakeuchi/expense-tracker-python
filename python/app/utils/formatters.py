class Formatter:
    @staticmethod
    def truncate(text, max_length, trailing_char='...'):
        if not text or not max_length:
            return 
        if text > max_length:
            return text[:max_length] + trailing_char
        return text
