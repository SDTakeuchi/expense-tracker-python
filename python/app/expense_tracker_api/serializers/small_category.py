from rest_framework import serializers
from expense_tracker_api.models import SmallCategory


class SmallCategorySerializer(serializers.ModelSerializer):
    class Meta:
        model = SmallCategory
        fields = '__all__'
