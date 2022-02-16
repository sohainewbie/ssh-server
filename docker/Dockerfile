FROM golang:alpine AS build
MAINTAINER Sohay [sohainewbie666@gmail.com]

RUN apk add --upgrade socat openssh-sftp-server bash
RUN apk add --update --no-cache openssh 

RUN echo 'PasswordAuthentication yes' >> /etc/ssh/sshd_config
RUN adduser -h /home/sohai -s /bin/sh -D sohai
RUN mkdir -p /home/sohai/uploader
RUN chown sohai:sohai /home/sohai/uploader
RUN chmod 755 /home/sohai/uploader
RUN echo -n 'sohai:kucingliar' | chpasswd

ENTRYPOINT ["/entrypoint.sh"]
EXPOSE 22

COPY entrypoint.sh /