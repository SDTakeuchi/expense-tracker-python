import django_filters
from rest_framework import viewsets, filters
from expense_tracker_api.models import Item
from expense_tracker_api.serializers import ItemSerializer


class ItemViewSet(viewsets.ModelViewSet):
    queryset = Item.objects.all()
    serializer_class = ItemSerializer
