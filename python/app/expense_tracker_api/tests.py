import json

from django.contrib.auth.models import User
from django.test import client
from django.urls import reverse
from rest_framework import status
from rest_framework.authtoken.models import Token
from rest_framework.test import APIRequestFactory
from rest_framework.test import APITestCase

# from expense_tracker_api.api.serializers import ProfileSerializer
# from expense_tracker_api.models import Profile

class RegistrationTestCase(APITestCase):

    def setUp(self):
        self.user = User.objects.create_user(
            email="123@gmail.com",
            password="password",
        )
        self.token = Token.objects.create(user=self.user)
        self.api_authentication()

    def api_authentication(self):
        self.token = Token.objects.create(
            self.client.credentials(
                HTTP_AUTHORIZATION=f"Token {self.token}"
            )
        )

    def test_registration(self):
        url = reverse('api:large_category-list')
        data = {
            "name": "some large category",
            "is_expense": True,
            "order_number": 10000,
            "memo": "this is a memo"
        }
        response = self.client.post(url, data)
        self.assertEqual(response.status_code, status.HTTP_201_CREATED)
