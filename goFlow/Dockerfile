FROM ubuntu:latest

LABEL name="newball"

WORKDIR  /usr/newball/github_demo/linux_v02

COPY ./conf/config.ini ./conf/config.ini

ADD ./goflow ./goflow

EXPOSE 8000

ENTRYPOINT ["./goflow","config.ini"]