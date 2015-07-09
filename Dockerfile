FROM golang:1.3.3

# install gpg2
RUN apt-get update && apt-get install -y gnupg2

#VOLUME /sign-server
#WORKDIR /sign-server

EXPOSE 8080 8081


ADD . /sign-server
ADD run-sign-server.sh /
RUN chmod u+x /run-sign-server.sh

CMD ["/run-sign-server.sh"]
