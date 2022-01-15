import json

from django.contrib.auth.models import User
from django.urls import reverse
from profiles.api.serializers import ProfileSerializer
from profiles.models import Profile
from rest_framework import status
from rest_framework.authentication.models import Token
from rest_framework.test import APITestCase
