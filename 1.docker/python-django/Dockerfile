FROM python:3.6.8-jessie

RUN pip install django

ADD . /app
ENTRYPOINT python /app/manage.py runserver 0.0.0.0:8000
