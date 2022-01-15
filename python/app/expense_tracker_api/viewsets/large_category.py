import django_filters
from expense_tracker_api.models import LargeCategory
from expense_tracker_api.serializers import LargeCategorySerializer
from rest_framework import filters, viewsets
from rest_framework.permissions import IsAuthenticated


class LargeCategoryViewSet(viewsets.ModelViewSet):
    queryset = LargeCategory.objects.all()
    serializer_class = LargeCategorySerializer
    permission_classes = [IsAuthenticated]
