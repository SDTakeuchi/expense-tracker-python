import django_filters
from expense_tracker_api.models import SmallCategory
from expense_tracker_api.serializers import SmallCategorySerializer
from rest_framework import filters, viewsets


class SmallCategoryViewSet(viewsets.ModelViewSet):
    queryset = SmallCategory.objects.all()
    serializer_class = SmallCategorySerializer
