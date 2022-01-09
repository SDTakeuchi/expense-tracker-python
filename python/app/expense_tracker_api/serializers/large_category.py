from rest_framework import serializers
from expense_tracker_api.models import LargeCategory


class LargeCategorySerializer(serializers.ModelSerializer):
    class Meta:
        model = LargeCategory
        fields = '__all__'
