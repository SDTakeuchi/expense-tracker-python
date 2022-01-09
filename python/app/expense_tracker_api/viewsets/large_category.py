import django_filters
from rest_framework import viewsets, filters
from expense_tracker_api.models import LargeCategory
from expense_tracker_api.serializers import LargeCategorySerializer


class LargeCategoryViewSet(viewsets.ModelViewSet):
    queryset = LargeCategory.objects.all()
    serializer_class = LargeCategorySerializer
