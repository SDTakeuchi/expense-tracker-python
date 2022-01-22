import uuid
from django.db import models


class BaseModel(models.Model):
    class Meta:
        abstract = True

    id = models.UUIDField(
        primary_key=True,
        default=uuid.uuid4,
        editable=False
    )

    created_at = models.DateTimeField(
        auto_now_add=True
    )

    updated_at = models.DateTimeField(
        auto_now=True
    )

    def get_or_none(self, **kwargs):
        """
        return obj if matching query exists
        else returns None
        """
        # maybe this has to be self.filter(kwargs).first()
        return self.objects.filter(kwargs).first()

    # def validate_same_value_exists(self, **kwargs):
    #     for k, v in kwargs.items():
