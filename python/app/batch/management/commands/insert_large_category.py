import sys

from django.core.management.base import BaseCommand
from django.db import transaction
from expense_tracker_api.models import LargeCategory
from tqdm import tqdm


class Command(BaseCommand):
    @transaction.atomic()
    def handle(self, *args, **options):
        self.insert_large_category()

    def insert_large_category(self):
        rows = self.get_large_category()
        print(sys._getframe().f_code.co_name)

        for row in tqdm(rows):
            try:
                name         = row[0]
                order_number = row[1]
                memo         = row[2] if row[2] else ''
                is_expense   = row[3] if row[3] else False 

                LargeCategory.objects.update_or_create(
                    name=name,
                    order_number=order_number,
                    memo=memo,
                    is_expense=is_expense
                )
            except Exception:
                raise Exception(f'An error occured in row {row}')

    def get_large_category(self):
        """
        in order of ...
        name, order_number, memo, is_expense
        """
        return (
            ('食費', 100,),
            ('趣味・娯楽', 200,),
            ('未分類', 99999,),
        )
