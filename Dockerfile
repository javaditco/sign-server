FROM golang:1.3.3

VOLUME /sign-server
WORKDIR /sign-server

EXPOSE 8080

ADD run-sign-server.sh /
RUN chmod u+x /run-sign-server.sh

CMD ["/run-sign-server.sh"]

