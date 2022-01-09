from django.contrib import admin
from django.urls import path, include
from rest_framework import routers
from .viewsets import *


app_name = 'api'

router = routers.DefaultRouter()
router.register(r'items', ItemViewSet)
router.register(r'large-categories', LargeCategoryViewSet)
router.register(r'small-categories', SmallCategoryViewSet)
router.register(r'users', UserViewSet)

urlpatterns = [
    path('', include(router.urls)),
]