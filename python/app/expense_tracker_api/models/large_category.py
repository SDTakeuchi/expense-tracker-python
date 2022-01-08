from django.db import models
from .user import CustomUser

class LargeCategory(models.Model):
    name = models.CharField(
        max_length=255
    )

    memo = models.CharField(
        max_length=255,
        blank=True,
        null=True
    )

    is_expense = models.BooleanField(
        default=True
    )

    order_number = models.IntegerField()

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
