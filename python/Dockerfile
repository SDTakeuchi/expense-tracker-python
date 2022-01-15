FROM python:3.10.1-slim

ENV PYTHONDONTWRITEBYTECODE 1
ENV PYTHONUNBUFFERED 1
ENV LANG C.UTF-8
ENV TZ Asia/Tokyo

RUN mkdir /app \
&&  mkdir /app/uwsgi \
&&  touch /app/uwsgi/uwsgi.log

COPY ./app/ /app/

WORKDIR /app/

RUN chmod 600 .env \
&&  apt-get update \
# for uwsgi not to fail
&&  apt-get -y install gcc \
# below line is needed to use opencv
# &&  apt-get install -y libgl1-mesa-glx libglib2.0-0 libsm6 libxrender1 libxext6 \
&&  pip3 install --upgrade pip \
&&  pip3 install -r requirements.txt

EXPOSE 8000

ENTRYPOINT ["uwsgi", "--ini", "/app/uwsgi.ini"]
