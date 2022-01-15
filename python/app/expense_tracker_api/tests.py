import json

from django.contrib.auth.models import User
from django.urls import reverse
from rest_framework import status
from rest_framework.authentication.models import Token
from rest_framework.test import APITestCase

from profiles.api.serializers import ProfileSerializer
from profiles.models import Profile

class RegistrationTestCase(APITestCase):

    def test_registration(self):
        data = {
            "email": "123@email.com",
            "password1": "strong-password",
            "password2": "strong-password",
        }
        response = self.client.post()
