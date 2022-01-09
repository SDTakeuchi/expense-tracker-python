import django_filters
from rest_framework import viewsets, filters
from expense_tracker_api.models import CustomUser
from expense_tracker_api.serializers import UserSerializer


class UserViewSet(viewsets.ModelViewSet):
    queryset = CustomUser.objects.all()
    serializer_class = UserSerializer
