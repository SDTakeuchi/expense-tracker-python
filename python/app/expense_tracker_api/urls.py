from django.contrib import admin
from django.urls import include, path
from rest_framework import routers

from .viewsets import *


app_name = 'api'

router = routers.DefaultRouter()
router.register(r'items', ItemViewSet)
router.register(r'large-categories', LargeCategoryViewSet, basename='large_category')
router.register(r'small-categories', SmallCategoryViewSet, basename='small_category')
router.register(r'users', UserViewSet)

urlpatterns = [
    path('', include(router.urls)),
]
