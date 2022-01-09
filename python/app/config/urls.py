from django.contrib import admin
from django.urls import path, include
from django.conf.urls import include
from expense_tracker_api.urls import router as api_router


urlpatterns = [
    # path('admin/', admin.site.urls),
    # re_path(r'^api/', include(api_router.urls)),
    path('admin/', admin.site.urls),
    path('api/', include('expense_tracker_api.urls')),
]
