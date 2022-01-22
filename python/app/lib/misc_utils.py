class MiscUtil:
    @staticmethod
    def get_current_url(request):
        return "{0}://{1}".format(request.scheme, request.get_host())

    @staticmethod
    def get_prev_url(request):
        if (prev_page := request.META.get('HTTP_REFERER')) is not None:
            return "{0}://{1}".format(prev_page.scheme, prev_page.get_host())
