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
        returns obj if one matching query exists
        else returns None
        if multiple objects are fecthed, wlil throw an error
        """
        try:
            return self.get_queryset().get(kwargs)
        except self.model.DoesNotExist:
            return None
        except Exception as e:
            raise e
        
