FROM python:3.10.1-slim

ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1

COPY ./app/ /app/

WORKDIR /app/

RUN chmod 600 .env \
&& apt-get update \
# below line is needed to use opencv
# &&  apt-get install -y libgl1-mesa-glx libglib2.0-0 libsm6 libxrender1 libxext6 \
&&  pip install --upgrade pip \
&&  pip install -r requirements.txt

EXPOSE 8000

ENTRYPOINT ["./manage.py", "runserver", "0.0.0.0:8000"]
