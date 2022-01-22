from django.db import models

from .base_model import BaseModel
from .large_category import LargeCategory
from .user import CustomUser


class SmallCategory(models.Model):
    name = models.CharField(
        max_length=255
    )

    memo = models.CharField(
        max_length=255,
        blank=True,
        null=True
    )

    order_number = models.IntegerField()

    large_category = models.ForeignKey(
        LargeCategory,
        on_delete=models.CASCADE,
        related_name="related_small_categories"
    )

    user = models.ForeignKey(
        CustomUser,
        on_delete=models.CASCADE,
        null=True,
        blank=True
    )

    def __str__(self):
        return self.name

    class Meta:
        ordering = ('order_number',)
