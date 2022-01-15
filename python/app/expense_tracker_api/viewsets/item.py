import django_filters
from django.utils.decorators import method_decorator
from django.views.decorators.cache import cache_page
from django.views.decorators.vary import vary_on_cookie
from expense_tracker_api.models import Item
from expense_tracker_api.serializers import ItemSerializer
from rest_framework import filters, viewsets


class ItemViewSet(viewsets.ModelViewSet):
    queryset = Item.objects.all()
    serializer_class = ItemSerializer

    @method_decorator(vary_on_cookie)
    @method_decorator(cache_page(60*60))
    def dispatch(self, *args, **kwargs):
        return super(ItemViewSet, self).dispatch(*args, **kwargs)
