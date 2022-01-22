from django.db import models

from .base_model import BaseModel
from .small_category import SmallCategory
from .user import CustomUser


class Item(BaseModel):
    name = models.CharField(
        max_length=255
    )

    memo = models.CharField(
        max_length=255,
        blank=True,
        null=True
    )

    price = models.IntegerField(
        default=0
    )

    purchase_date = models.DateTimeField()

    small_category = models.ForeignKey(
        SmallCategory,
        on_delete=models.SET_NULL,
        related_name="linked_items",
        null=True,
        blank=True
    )

    user = models.ForeignKey(
        CustomUser,
        on_delete=models.CASCADE
    )

    def __str__(self):
        return f'{self.name}({self.user}) - {self.purchase_date}'

    class Meta:
        ordering = ['-purchase_date']
