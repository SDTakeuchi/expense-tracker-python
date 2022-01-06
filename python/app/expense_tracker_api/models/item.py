import uuid
from django.db import models
from .small_category import SmallCategory
from .user import CustomUser

class LargeCategory(models.Model):
    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False
    )

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

    date_created = models.DateTimeField(
        auto_now_add=True
    )

    date_updated = models.DateTimeField(
        auto_now=True
    )

    small_category = models.ForeignKey(
        SmallCategory,
        on_delete=models.SET_NULL,
        related_name="linked_items"
    )

    user = models.ForeignKey(
        CustomUser,
        on_delete=models.CASCADE
    )

    def __str__(self):
        return self.name

    class Meta:
        ordering = ['-purchase_date']
