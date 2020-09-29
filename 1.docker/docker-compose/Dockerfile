FROM python:3.6.8-jessie

ENV PYTHONUNBUFFERED 1

WORKDIR /app

ADD requirements.txt .
RUN pip install -r requirements.txt

ADD . .

CMD /bin/sh docker-entrypoint.sh
