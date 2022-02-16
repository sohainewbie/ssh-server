FROM golang:alpine AS build
RUN apk add --upgrade socat openssh-sftp-server bash
RUN apk add --update --no-cache openssh 

RUN echo 'PasswordAuthentication yes' >> /etc/ssh/sshd_config
RUN adduser -h /home/sohai -s /bin/sh -D sohai
RUN echo -n 'sohai:kucingliar' | chpasswd

ENTRYPOINT ["/entrypoint.sh"]
EXPOSE 2222

COPY entrypoint.sh /
