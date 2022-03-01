class MiscUtil:
    @staticmethod
    def get_current_url(request):
        return "{0}://{1}".format(request.scheme, request.get_host())

    @staticmethod
    def get_prev_url(request):
        if (prev_page := request.META.get('HTTP_REFERER')) is not None:
            return "{0}://{1}".format(prev_page.scheme, prev_page.get_host())

    @staticmethod
    def cvt_char_bool_to_bool(val: str) -> bool:
        if val == 'True':
            return True
        elif val == 'False':
            return False
        return TypeError(
            "MiscUtil.cvt_char_bool_to_bool got neither 'True' or 'False'; "
            f"type: {type(val)}; "
            f"value: {val}"
        )

    @staticmethod
    def cvt_bool_to_char_bool(val: bool) -> str:
        if val is True:
            return 'True'
        elif val is False:
            return 'False'
        return TypeError(
            "MiscUtil.cvt_bool_to_char_bool got neither True or False; "
            f"type: {type(val)};  "
            f"value: {val}"
        )
