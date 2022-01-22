class Checker:
    @staticmethod
    def is_empty(value):
        """
        return True
        if value is any of
        None, '', [], {}
        """
        return value is None or value == '' or value == [] or value == {}

    @staticmethod
    def is_not_empty(value):
        """
        return False
        if value is any of
        None, '', [], {}
        """
        return value is not None and value != '' and value != [] and value != {}
