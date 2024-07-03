FROM postgres:latest

ENV DB_MASTER_USER=postgres
ENV DB_MASTER_PASS=password
ENV DB_HOST=postgres
ENV DB_PORT=5432
ENV DB_NAME=salesorder

RUN apt-get update && apt-get install -y gettext


RUN echo  "permissionset ronald"

COPY ./init.sh /docker-entrypoint-initdb.d/init.sh
