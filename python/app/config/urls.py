from django.contrib import admin
from django.urls import path, include


urlpatterns = [
    path('api/admin/', admin.site.urls),
    path('api/auth/', include('djoser.urls')),
    path('api/auth/', include('djoser.urls.jwt')),
    path('api/v1/', include('expense_tracker_api.urls')),
]
