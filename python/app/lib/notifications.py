from django.core.mail import EmailMessage
from config.settings import EMAIL_HOST_USER


class Notifications:
    @staticmethod
    def send_email(
        recipients: list,
        subject,
        body,
        bcc_enabled=True,
    ):
        if not isinstance(recipients, list):
            raise ValueError("recipient for Notifications.send_email must be either list.")
        if bcc_enabled is True:
            email = EmailMessage(subject, body, EMAIL_HOST_USER, bcc=recipients)
        else:
            email = EmailMessage(subject, body, EMAIL_HOST_USER, to=recipients)
        email.send()

    @staticmethod
    def create_user_notification(
        recipient,
        subject,
        body
    ):
        pass
