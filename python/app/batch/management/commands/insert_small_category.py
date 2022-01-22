import sys

from django.core.management.base import BaseCommand
from django.db import transaction
from expense_tracker_api.models import LargeCategory, SmallCategory
from tqdm import tqdm


class Command(BaseCommand):
    @transaction.atomic()
    def handle(self, *args, **options):
        self.insert_small_category()

    def insert_small_category(self):
        rows = self.get_small_category()
        print(sys._getframe().f_code.co_name)

        for row in tqdm(rows):
            try:
                large_category_name = row[0]
                name                = row[1]
                order_number        = row[2]
                memo                = row[3] if row[3] else ''

                large_category = LargeCategory.objects.get_or_none(
                        name=large_category_name
                    )
                if large_category is None:
                    raise Exception('something went wrong with codes in small ctegory batch.')

                SmallCategory.objects.update_or_create(
                    large_category=large_category,
                    name=name,
                    order_number=order_number,
                    memo=memo
                )
            except Exception:
                raise Exception(f'An error occured in row {row}')

    def get_small_category(self):
        """
        in order of ...
        large_category, name, order_number, memo
        """
        return (
            ('食費', '食料品', 100, ),
            ('食費', '外食', 200,),
        )
