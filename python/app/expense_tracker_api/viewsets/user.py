import django_filters
from expense_tracker_api.models import CustomUser
from expense_tracker_api.serializers import UserSerializer
from rest_framework import filters, viewsets


class UserViewSet(viewsets.ModelViewSet):
    queryset = CustomUser.objects.all()
    serializer_class = UserSerializer
