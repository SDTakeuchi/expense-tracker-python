import django_filters
from expense_tracker_api.models import CustomUser
from expense_tracker_api.serializers import UserSerializer
from rest_framework import filters, viewsets
from rest_framework.permissions import IsAuthenticated


class UserViewSet(viewsets.ModelViewSet):
    permission_classes = [IsAuthenticated]
    queryset = CustomUser.objects.all()
    serializer_class = UserSerializer
