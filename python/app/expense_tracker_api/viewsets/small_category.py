import django_filters
from rest_framework import viewsets, filters
from expense_tracker_api.models import SmallCategory
from expense_tracker_api.serializers import SmallCategorySerializer


class SmallCategoryViewSet(viewsets.ModelViewSet):
    queryset = SmallCategory.objects.all()
    serializer_class = SmallCategorySerializer
