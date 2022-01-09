from django.contrib import admin
from .models import *


admin.site.register(CustomUser)
admin.site.register(LargeCategory)
admin.site.register(SmallCategory)
admin.site.register(Item)
