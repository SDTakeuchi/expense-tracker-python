from batch.management.commands import (
    insert_large_category,
    insert_small_category,
)
from django.core.management.base import BaseCommand


class Command(BaseCommand):
    """
    executes all the batches for initial setup
    """
    def handle(self, *args, **options):
        print('initital set up started...')

        insert_large_category.Command().handle()
        insert_small_category.Command().handle()

        print('initial set up done.')
