FROM busybox
MAINTAINER Gustavo Gimenez <gimenezanderson@gmail.com>
RUN mkdir /app
ADD build/server /app
RUN mkdir images
ADD images /images
RUN chmod -R 777 /app
EXPOSE 9000

ENTRYPOINT "/app/server"
