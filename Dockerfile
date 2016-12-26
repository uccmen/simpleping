FROM gliderlabs/alpine:latest

MAINTAINER Ucchishta Sivaguru <ucc.men@gmail.com>

RUN apk-install git make py-pip go && pip install awscli

RUN mkdir -p /opt/go && \
    mkdir -p /home/app && \
    adduser -D app

WORKDIR /home/app
ENV GOPATH /opt/go

COPY . /home/app

RUN chown -R app:app /home/app && \
    make deps && make

USER app

EXPOSE 3000
ENTRYPOINT ["make"]
CMD ["run"]